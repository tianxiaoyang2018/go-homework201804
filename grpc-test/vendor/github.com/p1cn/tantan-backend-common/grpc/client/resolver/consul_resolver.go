package resolver

// type ConsulResolver struct {
// 	serviceName string
// 	tags        []string
// }

// func NewConsulResolver(serviceName string, tags []string) (naming.Resolver, error) {
// 	return &ConsulResolver{
// 		serviceName: serviceName,
// 		tags:        tags,
// 	}, nil
// }

// // target example: "http://127.0.0.1:8888"
// func (self *ConsulResolver) Resolve(target string) (naming.Watcher, error) {

// 	dcfg := api.DefaultConfig()
// 	dcfg.Address = target

// 	apiClient, err := api.NewClient(dcfg)
// 	if err != nil {
// 		return nil, err
// 	}
// 	client := consul.NewClient(apiClient)

// 	return newConsuleWatcher(client, self.serviceName, self.tags)
// }

// func newConsuleWatcher(client consul.Client, serviceName string, tags []string) (naming.Watcher, error) {
// 	logger := log.NewLogfmtLogger(os.Stderr)
// 	watcher := consul.NewWatcher(client, logger, serviceName, tags, true)

// 	notify := make(chan sd.Event, 10)
// 	watcher.Register(notify)

// 	return &consuleWatcher{
// 		client:      client,
// 		serviceName: serviceName,
// 		tags:        tags,
// 		watcher:        watcher,
// 		notifyChan:  notify,
// 	}, nil

// }

// type consuleWatcher struct {
// 	client      consul.Client
// 	serviceName string
// 	tags        []string
// 	watcher        *consul.Watcher
// 	notifyChan  chan sd.Event
// 	addresses   map[string]struct{}
// }

// func (self *consuleWatcher) Close() {
// 	self.watcher.Stop()
// }

// func (self *consuleWatcher) Next() ([]*naming.Update, error) {

// 	event := <-self.notifyChan
// 	if event.Err != nil {
// 		return nil, event.Err
// 	}

// 	var updates []*naming.Update

// 	for k, _ := range self.addresses {
// 		found := false
// 		for _, address := range event.Instances {
// 			if k == address {
// 				found = true
// 				break
// 			}
// 		}
// 		if !found {
// 			updates = append(updates, &naming.Update{
// 				Op:   naming.Delete,
// 				Addr: k,
// 			})
// 		}
// 	}

// 	for _, address := range event.Instances {
// 		if _, ok := self.addresses[address]; !ok {
// 			updates = append(updates, &naming.Update{
// 				Op:   naming.Add,
// 				Addr: address,
// 			})
// 			self.addresses[address] = struct{}{}
// 		}
// 	}

// 	return updates, nil
// }
