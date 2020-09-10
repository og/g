#!/bin/bash
cd conv && go test &&  cd ..
cd crypto && go test &&  cd ..
cd error && go test &&  cd ..
cd json && go test &&  cd ..
cd map && go test &&  cd ..
cd math && go test &&  cd ..
cd rand && go test &&  cd ..
cd reject && go test &&  cd ..
cd test && go test &&  cd ..
cd time && go test &&  cd ..