1.package main
import "fmt"

func enqueue(queue[] int, element int) []int {
  queue = append(queue, element); // Simply append to enqueue.
  fmt.Println("Enqueued:", element);
  return queue
}

func dequeue(queue[] int) ([]int) {
  element := queue[0]; // The first element is the one to be dequeued.
  fmt.Println("Dequeued:", element)
  return queue[1:]; // Slice off the element once it is dequeued.
}

func main() {
  var queue[] int; // Make a queue of ints.

  queue = enqueue(queue, 100);
  queue = enqueue(queue, 200);
  queue = enqueue(queue, 300);

  fmt.Println("Queue:", queue);

  queue = dequeue(queue);
  fmt.Println("Queue:", queue);

  queue = enqueue(queue, 400);
  fmt.Println("Queue:", queue);
}
Enqueued: 100
Enqueued: 200
Enqueued: 300
Queue: [100 200 300]