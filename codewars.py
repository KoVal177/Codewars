# -*- coding: utf-8 -*-
def xo(s):

    numX = 0
    numO = 0
    VarEx = True
    s = s.lower()
    for i in s:
        if (i == "x"):
            numX += 1
        elif (i == "o"):
            numO += 1
    
    if (numX == numO) :
        VarEx = True
    else :
        VarEx = False
    return VarEx


def row_sum_odd_numbers(n):

    PrevNums = n*(n-1)/2                    # calculating number of numbers in previous rows (progression)
    NumStart = PrevNums*2+1                 # calculating first number in n-th row
    NumEnd = NumStart+(n-1)*2               # calculating last number in n-th row
    SumOdd = int((NumStart+NumEnd)*n/2)     # calculating sum of numbers in n-th row (progression) and put it into return variable
    
    return SumOdd


def nb_year(p0, percent, aug, p):
    pEnd = p0
    years = 0
    percent = percent/100
    while pEnd < p :
        pEnd = int(pEnd + pEnd*percent + aug)
        years += 1
    return years

def solution(number):
    Res = 0
    if number < 0: return Res
    for i in range(number):
        if i%3 == 0 :
            Res = Res+i
        elif i%5 == 0 :
            Res = Res+i
    
    return Res

def to_jaden_case(string):
    strEnd = ""
    for i in string.split(' '):
        strEnd = strEnd + i.capitalize() + " "
    return strEnd.strip()


def digital_root(n):
    Res=n
    while Res >= 10 :
        n = Res
        Res = 0
        while n > 0:
            Res = Res+n%10
            n = n//10
    
    return Res

def divisors(integer):
    resArray = []
    for i in range(2,integer):
        if integer%i == 0:
            resArray.append(i)
    if not(resArray) :
        return (f"{integer} is prime")
    return resArray

def array_diff(a, b):
    for i in b :
        while i in a :
            a.remove(i)
    return a

def move_zeros(array):
    for i in array :
        if i == 0 :
            array.remove(i)
            array.append(i)
    return array

def rot13(message):
    InAphbt = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
    OutAphbt = "nopqrstuvwxyzabcdefghijklmNOPQRSTUVWXYZABCDEFGHIJKLM"
    mesOut = ""
    for i in message :
        num = InAphbt.find(i)
        if num >= 0:
            mesOut = mesOut + OutAphbt[num]
        else:
            mesOut = mesOut + i
    return mesOut

def find_outlier(integers):
    test = integers[0]%2 + integers[1]%2 + integers[2]%2
    if test < 2 :
        test = 1
    else:
        test = 0
    for i in integers :
        if i%2 == test :
            return i
    
    return None

def solution(string,markers):
    OutStr = ""
    tmp = string.split("\n")
    for i in tmp:
        for j in markers :
            i = i.split(j)[0].strip()
        OutStr = OutStr + i + "\n"
    return OutStr[:len(OutStr)-1]

def generate_hashtag(s):
    if s == "": return False
    strOut = "#"
    tmp = s.split()
    for i in tmp :
        strOut = strOut + i.capitalize()
    if len(strOut) > 140: return False
    return strOut

def multipl_arrs (a, b): # multiplie arrays
    c = []
    for i in a :
        for j in b :
            c.append(i + j)
    return (c)

def get_pins(observed):
    possibilities = ("08", "124", "2135", "326", "4157", "52468", "6359", "748", "85790", "968")
    PsblOut = []    
    for i in possibilities[int(observed[0])] : # initialization with first digit
        PsblOut.append(i)
        
    for i in range(1, len(observed)) :          # adding next digits to possible pins
        PsblOut = multipl_arrs(PsblOut, possibilities[int(observed[i])])
    return PsblOut

def score(dice):
    ResScore = 0
    Matrix = [0, 0, 0, 0, 0, 0]
    Score3 = [1000, 200, 300, 400, 500, 600]
    Score1 = [100, 0, 0, 0, 50, 0]
    for i in dice :
        Matrix[i-1] += 1
    for i in range(0, 6) :
        ResScore += Score3[i]*int(Matrix[i]//3) + Score1[i]*int(Matrix[i]%3)
    return ResScore

def solution(args):
    RangeStr = ""
    Last = args[0]
    FlagPoss = False           # flag if local range is possible (2 consecutive)
    FlagDef = False            # flag if local range is definite (at least 3 consecutive)
    for i in args:
        if (i == Last+1):      # check if last and current are consecutive numbers
            if FlagPoss :          # were 2 consecutives already so at least 3, change flag
                FlagDef = True
            else :                  # only first time consecutive - raise possible flag
                FlagPoss = True
            Last = i
        elif FlagDef :              # no more sonsecutive, but it was local rang
            RangeStr += "-" + str(Last) + "," + str(i)
            FlagDef = False;
            FlagPoss = False;
            Last = i
        elif FlagPoss :             # only 2 consecutive, no local range
            RangeStr += "," + str(Last) + "," + str(i)
            FlagDef = False;
            FlagPoss = False;
            Last = i
        else:                       # no consecutive
            RangeStr += "," + str(i)
            FlagDef = False;
            FlagPoss = False;
            Last = i
    if FlagDef :                    # final additions for not handled in cycle
        RangeStr += "-" + str(Last)
    elif FlagPoss :
        RangeStr += "," + str(Last)
    RangeStr = RangeStr[1:]      # remove 1st comma
    return RangeStr


def permutations(string):
    def add_elem_perm(permIn, char):
        permOut = set()
        for i in range(len(permIn[0])+1) :
            for j in permIn :
                tmp = j[0:i] + char + j[i:]
                permOut.add(tmp)    
        return list(permOut)

    perms = string[0]        
    for i in string[1:]:
        perms = add_elem_perm(perms, i)
    return perms

def duplicate_encode(word):
    EncWord = ""
    word = word.lower()
    for i in word :
        if word.count(i) > 1 :
            EncWord += ")"
        else:
            EncWord += "("
    
    return EncWord

def top_3_words(text):
    AcceptedChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ'"
    words = {}
    sorted_words = []
    text = text.lower()
    textNew = ""
    for i in text :
        if i in AcceptedChars :
            textNew += i
        else :
            textNew += " "
    tmp = textNew.split()
    for i in tmp :
        if i.replace("'","") != "" :
            if i in words :
                words[i] += 1
            else :
                words.update({i:1})

    sorted_keys = sorted(words, key=words.get, reverse=True)
    for i in sorted_keys[:3] :
        sorted_words.append(i)
    
    return sorted_words


def hamming_old(n):
    hamms = set()
    hamms.add(1)
    i = 1
    numRes = 1
    while i < n :
        numRes += 1
        if (numRes%2 == 0) and (int(numRes//2) in hamms) : 
            i += 1
            hamms.add(numRes)
        elif (numRes%3 == 0) and (int(numRes//3) in hamms) : 
            i += 1
            hamms.add(numRes)
        elif (numRes%5 == 0) and (int(numRes//5) in hamms) : 
            i += 1
            hamms.add(numRes)
            
    return numRes

def hamming(n):    
    hamms = [1]
    i = 1
    i2 = i3 = i5 = 0
    Last = 1
    while i < n :
        Last = min(hamms[i2]*2, hamms[i3]*3, hamms[i5]*5)
        hamms.append(Last)
        if Last%2 == 0 : i2 += 1
        if Last%3 == 0 : i3 += 1
        if Last%5 == 0 : i5 += 1
        i +=1
    return Last

def decode_rail_fence_cipher(string, n):
    output=""
    if string == "" : return output
    
    # create pattern to move through
    pattern = []
    pattlength = n*2-2
    for i in range(n):
        pattern.append(i)
    for i in range(n-2,0,-1):
        pattern.append(i)

    # count lengths of the strings
    # lengths if round number of patterns
    length = [len(string)//pattlength*2] * n
    length[0] = length[n-1] = len(string)//(pattlength)
    # extra length for the rest
    for i in range(len(string)%(pattlength)):
        length[pattern[i]] += 1

    # initialize starting points of each string
    index = [0] * n
    for i in range(1,n):
        index[i] = index[i-1] + length[i-1]
    
    # walk through the string with pattern
    for i in range(len(string)//pattlength+1):
        for j in pattern:
            if len(output) ==  len(string) : break
            output += string[index[j]]
            index[j] += 1
    
    return output

def encode_rail_fence_cipher(string, n):
    output=""
    if string == "" : return output

    # create pattern to move through
    pattern = []
    pattlength = n*2-2
    for i in range(n):
        pattern.append(i)
    for i in range(n-2,0,-1):
        pattern.append(i)

    # walk through the string with pattern
    parts = [""] * n
    i = 0
    for k in range(len(string)//pattlength+1):
        for j in pattern:
            if i >= len(string) : break
            parts[j] += string[i]
            i += 1
    
    output = "".join(parts)
    return output

class Node:
    def __init__(self, L, R, n):
        self.left = L
        self.right = R
        self.value = n

def tree_by_levels(node):
    output = []
    curr = []
    if node == None : return output
    
    last = start = 0
    curr.append(node)
    while True :
        if start != len(curr) :
            last = len(curr)
            for i in range(start, len(curr)) :
                if curr[i].left != None :
                    curr.append(curr[i].left)
                if curr[i].right != None :
                    curr.append(curr[i].right)
            start = last
        else: 
            break
            
    for i in range(len(curr)) :
        output.append(curr[i].value)
    return output

def next_direction(old_dir) :
    if   old_dir == [0,1] : new_dir = [1,0]
    elif old_dir == [1,0] : new_dir = [0,-1]
    elif old_dir == [0,-1] : new_dir = [-1,0]
    elif old_dir == [-1,0] : new_dir = [0,1]
    return new_dir

def snail(snail_map):
    if snail_map == [[]] : return []
    if len(snail_map) == 1 : return [snail_map[0][0]]
    border = len(snail_map)
    visited=[False] * border
    for i in range(border) : visited[i] = [False] * border

    coords = list([0, 0])
    direction = [0,1]
    newcoords = [0,0]
    
    output = []
    output.append(snail_map[coords[0]][coords[1]])
    visited[0][0] = True
    
    flag = 0
    while flag < 2 :
        newcoords = [x+y for x,y in zip(coords, direction)]
        if (-1< newcoords[0] <border) and (-1< newcoords[1] <border) :
            if not visited[newcoords[0]][newcoords[1]]:
                coords = newcoords
                visited[coords[0]][coords[1]] = True
                output.append(snail_map[coords[0]][coords[1]])
                flag = 0
            else:
                direction = next_direction(direction)
                flag += 1
        else:
            direction = next_direction(direction)
    return output

def next_smaller(n):
    if n//10 == 0 : return -1 # special case:one digit
    string = str(n)
    length = len(str(n))
    output = -1
    
    # moving from the end of the number to its head and looking for digit that can me changed for smaller on in the tail
    for i in range(length-2,-1,-1):
        maxmin = -1
        #checking tail for better numbers than current digit i
        for j in range(length-i-1):
            if int(string[i]) > int(string[i+j+1]) :
                # at least one digit found! still check the rest
                if maxmin < int(string[i+j+1]) :
                    ind_maxmin = i+j+1
                    maxmin = int(string[i+j+1])
        # if better digit found prepare the output and finish function
        if maxmin > -1 :
            beginning = string[:i] + str(maxmin)
            if beginning[0] == "0" : return -1     # special case: resulting number starts with "0", i.e. no possible outcomes
            ending = string[i+1:ind_maxmin] + string[i] + string[ind_maxmin+1:]
            output = int(beginning + "".join(sorted(ending,reverse=True)))
            return output
    return output


















