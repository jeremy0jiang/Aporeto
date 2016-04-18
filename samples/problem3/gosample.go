/**
 *@author Shenghao Jiang
 *@since 04-15-2016
 *@version 1.0
 */

package main

/**
 *This program will read muptiple urls entered in the command line
 *and generate word frequency table. Output will be stored in files
 *with name url1.txt,url2.txt...
 */
import (
		"strconv"
		"flag"
    "fmt"
    "net/http"
    "io/ioutil"
    "os"
		"strings"
		"regexp"
		"bufio"
    )

/**
 *method geturl
 *This method copys contents form the giving url
 *@param url
 *@return text
 */
func geturl(elements string) string {
    response,err := http.Get(elements)   
		if err != nil {
        fmt.Printf("%s", err)
        os.Exit(1)
    } else {
        defer response.Body.Close()
        contents, err := ioutil.ReadAll(response.Body)
        if err != nil {
            fmt.Printf("%s", err)
            os.Exit(1)
        }
		var  text string = string(contents)
    	return text
		}
	return "yes"
}

/**
 *method get_words_from
 *This method split text into words
 *@param text
 *@return words
 */ 
func get_words_from(text string) []string{
    words:= regexp.MustCompile("\\w+")
    return words.FindAllString(text, -1)
}

/**
 *method count_words
 *This method count the number of words in a giving string
 *@param words
 *@return word_counts
 */
func count_words (words []string) map[string]int{
    word_counts := make(map[string]int)
    for _, word :=range words{
        word_counts[word]++
    }
    return word_counts;
}

/**
 *method writetofile
 *This method opens and writes text to a file, where filename is 
 *passed as parameter
 *@param word_counts
 *@param path
 */
func writetofile (word_counts map[string]int,path string) error{
	file, err := os.Create(path)
  if err != nil {
    return err
  }
  defer file.Close()
   
	w := bufio.NewWriter(file)

  for word, word_count :=range word_counts{
  	fmt.Fprintln(w,word,":", word_count)

	}
	return w.Flush()
}



/**
 *method main
 *This method calls methods as described above and perform the task
 */
func main(){
		var urlname string
    flag.StringVar(&urlname, "urls", "common seperate one or more urls", "urls")
		flag.Parse()
	
		//var url string
		url := strings.Split(urlname,",")
    
    //loop through all entered urls by len(url) times
		for i :=0; i < len(url);i++{

		//read text from url and store text in result
 		result := geturl(url[i])
		
    //convert int i to string by the convenience to assign filename
    count := strconv.Itoa(i+1)
		
		//creat file name with url1.txt, url2.txt
		var filename string = "url"+count+".txt"
		
		//write word frequency table to output file
    writetofile(count_words(get_words_from(result)),filename)
		}
}
