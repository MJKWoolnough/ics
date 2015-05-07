package ics

type Calendar struct{}

func (c *Calendar) SetProductID(id string) {

}

type Event struct{}

func (c *Calendar) AddEvent(e *Event) {

}

type Todo struct{}

func (c *Calendar) AddTodo(t *Todo) {

}

type Journal struct{}

func (c *Calendar) AddJournal(j *Journal) {

}

type FreeBusy struct{}

func (c *Calendar) AddFreeBusy(fb *FreeBusy) {

}

type Timezone struct{}

func (c *Calendar) AddTimezone(tz *Timezone) {

}
