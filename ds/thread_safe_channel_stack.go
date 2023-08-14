package ds

type stackOperation int

var pushOp stackOperation = 1
var popOp stackOperation = 2
var emptyOp stackOperation = 3
var peekOp stackOperation = 4

type result struct {
	val interface{}
	err error
}

type request[T any] struct {
	op       stackOperation
	val      T
	response chan result
}

type ThreadSafeChannelStack[T any] struct {
	Stack[T]
	requestCh chan request[T]
}

var defaultThreadSafeStackSize = 0

func NewThreadSafeStack[T any](s Stack[T]) *ThreadSafeChannelStack[T] {
	ts := &ThreadSafeChannelStack[T]{
		Stack:     s,
		requestCh: make(chan request[T]),
	}
	go ts.serve()
	return ts
}

func (s *ThreadSafeChannelStack[T]) serve() {
	for req := range s.requestCh {
		switch req.op {
		case pushOp:
			s.Stack.Push(req.val)
		case peekOp:
			topEle, err := s.Stack.Peek()
			req.response <- result{
				val: topEle,
				err: err,
			}
		case emptyOp:
			empty := s.Stack.Empty()
			req.response <- result{
				val: empty,
				err: nil,
			}
		case popOp:
			topEle, err := s.Stack.Pop()
			req.response <- result{
				val: topEle,
				err: err,
			}
		default:
			panic("unknown operation")
		}
	}
}

func (s *ThreadSafeChannelStack[T]) sendRequest(op stackOperation, val T) (response chan result) {
	response = make(chan result)
	s.requestCh <- request[T]{
		op:       op,
		val:      val,
		response: response,
	}
	return response
}

func (s *ThreadSafeChannelStack[T]) Push(val T) {
	_ = s.sendRequest(pushOp, val)
}

func (s *ThreadSafeChannelStack[T]) Peek() (T, error) {
	var zeroVal T
	response := s.sendRequest(peekOp, zeroVal)
	res := <-response
	return res.val.(T), res.err
}

func (s *ThreadSafeChannelStack[T]) Pop() (T, error) {
	var zeroVal T
	response := s.sendRequest(popOp, zeroVal)
	res := <-response
	return res.val.(T), res.err
}

func (s *ThreadSafeChannelStack[T]) Empty() bool {
	var zeroVal T
	response := s.sendRequest(emptyOp, zeroVal)
	res := <-response
	return res.val.(bool)
}
