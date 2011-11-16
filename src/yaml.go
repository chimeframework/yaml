package yaml

import (
    "bytes"
    "fmt"
    "launchpad.net/goyaml"
    "os"
)

func Parse(path string) map[interface{}]interface{}{
   if data, err := ReadFileAsBytes(path); err != nil{
       panic(fmt.Sprintf("Error while reading YAML file %v", path))
   }

   output := make[map[interface{}]interface{}]

   if err := goyaml.Unmarshal(data, &output); err != nil{
       panic(err)
   }

   return output
}

func ReadFileAsBytes(filename string) (contents []byte, err error) {
    file, e := os.OpenFile(filename, os.O_RDONLY, 0644)
    if e != nil {
        return nil, e
    }

    var buf bytes.Buffer
    buf.ReadFrom(file)
    return buf.Bytes(), nil
}