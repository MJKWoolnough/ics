package ics

const vTodo = "VTODO"

type Todo struct {
}

func (c *Calendar) decodeTodo(d Decoder) error {
	for {
		p, err := d.p.GetProperty()
		switch p := p.(type) {
		case begin:
			if err = d.readUnknownComponent(string(p)); err != nil {
				return err
			}
		case end:
			if p != vTodo {
				return ErrInvalidEnd
			}
			return nil
		}
	}
	return nil
}
