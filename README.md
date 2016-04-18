#Thoughts and concerns about the projects

##Problem 1
1.In bash shell script, I read command line flags and output usages in my own way.That is, in problem 2 and 3, I read flags use go's built-in flags class, which is more convenient.

2.I do not quite understand why we need prompt "File removed", since we only need to overwrite file if it exists.

3.I am not sure if I can use file 50states and cat its content to a new file. If this is not allowed, I would a create long string array with names of 50 states, but which is not concise .

##Problem 2
1.The time complexity for my program is O(n),where readlines, remove, write each takes O(n). In terms of big O, I think my program is the fastest, which takes O(n). Readlines and write at least take O(n) each to read and write. In the part of remove, I used golang's built-in maps, which is similar to hash table. In my remove, it takes O(n) to create a map and it takes O(n) to remove duplicate from map. 

In terms of real time performance, I could have done better. I can combine read, remove, and write by directly reading to maps, removing from maps, and writing to file from maps instead of using reading string and then assining to maps.
 
2.This program looks like hard to me since I have never used go programming language before, but I tried to break the task down into small steps.
-read command line
-read file
-find duplicate
-write file

3.I learned the basic golang in a few hours and writing these code did not take me  very long. I think golang is very easy to read and easy to write.

##Problem 3
1.I have some concerns about this problem. I can successfully run the program on the url provided on github, but when I run this on other website, some html code will count towards word frequency.

2.I think I can solve this problem but it may take a few days. I am new to golang but I would like to learn new programming languages and accept new challenges.
