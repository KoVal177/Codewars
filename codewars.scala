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

object Sol {

  def evenOrOdd(number: Int): String =
    number % 2 match
      case 0 => "Even"
      case _ => "Odd"
}

def maskify(cc: String): String =
  if cc.length <= 4 then cc
  else "#" + maskify(cc.tail)

object MultiplesOf3Or5 {   
  def solution(number: Int): Long =
    def iter(threes: Int, fives: Int, s: Long, n: Int): Long =
      if threes + 3 >= n && fives + 5 >= n then s
      else if threes + 3 == fives + 5 then iter(threes + 3, fives + 5, s + threes + 3, n)
      else if threes + 3 < fives + 5 then iter(threes + 3, fives, s + threes + 3, n)
      else iter(threes, fives + 5, s + fives + 5, n)
    iter(0, 0, 0, number)
}

object WhichAreIn {

  def inArray(array1: Array[String], array2: Array[String]): Array[String] =
    def iter(a1: Array[String], a_res: Array[String], a_condition: String): Array[String] =
      if a1.isEmpty then a_res.sortWith(_ < _)
      else if (a_condition contains a1.head) && !(a_res contains a1.head) then 
        iter(a1.tail, a_res :+ a1.head, a_condition)
      else
        iter(a1.tail, a_res, a_condition)
    iter(array1, Array(), array2.mkString(" "))
}
