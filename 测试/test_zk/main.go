package main

import (
	"github.com/samuel/go-zookeeper/zk"
	"log"
	"sort"
	"strconv"
	"sync"
	"time"
)

func main() {
	//hosts := []string{"127.0.0.1:2182", "127.0.0.1:2183", "127.0.0.1:2184", "127.0.0.1:2185", "127.0.0.1:2186"}
	hosts := []string{"127.0.0.1:2182", "127.0.0.1:2185", "127.0.0.1:2186", "127.0.0.1:2183", "127.0.0.1:2184"}
	conn, _, err := zk.Connect(hosts, time.Second*5, zk.WithEventCallback(handleEvent))
	if err != nil {
		log.Fatalf("zk.Connect failed, err=%v", err)
	}
	defer conn.Close()


	//leadership.NewCandidate()

	//acls := zk.WorldACL(zk.PermAll)
	//exists, _, err := conn.Exists("/test")
	//if err != nil {
	//	log.Fatalf("conn.Exists failed, err=%v", err)
	//}
	//if !exists {
	//	res, err := conn.Create("/test", []byte("test"), 0, acls)
	//	if err != nil {
	//		log.Fatalf("conn.Create failed, err=%v", err)
	//	}
	//	log.Printf("conn.Create result=%s", res)
	//}
	//
	//wg := &sync.WaitGroup{}
	//for i := 0; i < 10; i++ {
	//	go lock(uint64(i), conn, wg)
	//}
	//wg.Wait()

	//for {
	//	data, _, err := conn.Get("/test")
	//	log.Printf("conn.Get, data=%s, err=%v", string(data), err)
	//	time.Sleep(time.Millisecond * 500)
	//}

	acls := zk.WorldACL(zk.PermAll)

	res, err := conn.Create("/test/test-ephemeral", []byte("111"), 1, acls)
	if err != nil {
		log.Fatalf("conn.CreateProtectedEphemeralSequential failed, err=%v", err)
	}
	log.Printf("res=%s", res)

	time.Sleep(time.Second * 300)

	//data, _, eventCh, err := conn.GetW("/test/test-ephemeral")
	//if err != nil {
	//	log.Fatalf("conn.GetW failed, err=%v", err)
	//}
	//log.Printf("data=%s", string(data))
	//select {
	//case event := <-eventCh:
	//	log.Printf("event=%v", event)
	//}
}

func handleEvent(event zk.Event) {
	log.Printf("handleEvent event=%v", event)
}

func lock(i uint64, conn *zk.Conn, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	acls := zk.WorldACL(zk.PermAll)

	res, err := conn.Create("/test/lock/lock", []byte(strconv.FormatUint(i, 10)), 3, acls)
	if err != nil {
		log.Fatalf("conn.CreateProtectedEphemeralSequential failed, err=%v, i=%d", err, i)
	}
	log.Printf("i=%d, res=%s", i, res)

	childs, _, err := conn.Children("/test/lock")
	if err != nil {
		log.Fatalf("conn.Children failed, err=%v, i=%d", err, i)
	}
	sort.Strings(childs)
	log.Printf("i=%d, childs=%v", i, childs)

	if "/test/lock/"+childs[0] != res {
		exists := false
		var eventCh <-chan zk.Event
		var err error
		for index, child := range childs {
			if res == "/test/lock/"+child {
				watchPath := "/test/lock/" + childs[index-1]
				log.Printf("i=%d, conn.ExistsW, path=%s", i, watchPath)
				exists, _, eventCh, err = conn.ExistsW(watchPath)
				if err != nil {
					log.Fatalf("conn.ExistsW failed, err=%v, i=%d", err, i)
				}
				break
			}
		}

		if exists {
			select {
			case event := <-eventCh:
				log.Printf("i=%d, event=%v", i, event)
			}
		}
	}

	log.Printf("i=%d, now is leader", i)
	time.Sleep(time.Second)
	log.Printf("i=%d, exit", i)
	_, stat, _ := conn.Get(res)
	_ = conn.Delete(res, stat.Version)
}
