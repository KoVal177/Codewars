object SheepAdvisor {
  def warnTheSheepCheck(queue: Array[String]): String = {
    def iter(queue: Array[String], i: Int): String = {
      if (queue.head == "wolf")
        s"Oi! Sheep number $i! You are about to be eaten by a wolf!"
      else
        iter(queue.tail, i-1)
    }

    if (queue.last == "wolf")
      "Pls go away and stop eating my sheep"
    else
      iter(queue, queue.length-1)
  }
}

