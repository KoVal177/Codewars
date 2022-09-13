object SheepAdvisor {
  def warnTheSheepCheck(queue: Array[String]): String = {
    def iter_func(queue: Array[String], i: Int): String = {
      if (queue.head == "wolf")
        s"Oi! Sheep number $i! You are about to be eaten by a wolf!"
      else
        iter_func(queue.tail, i-1)
    }

    if (queue.last == "wolf")
      "Pls go away and stop eating my sheep"
    else
      iter_func(queue, queue.length-1)
  }
}

