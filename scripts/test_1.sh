#!/bin/bash
for i in {1..5000}
do
    curl  http://localhost:3355/about
    curl  http://localhost:3355/
done


