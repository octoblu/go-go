#!/bin/bash

generate_templates(){
  go-bindata -prefix templates templates/...
}

build(){ 
  go build
}

main(){
  generate_templates || panic "generate_templates failed"
  build              || panic "build failed"
}
main $@
