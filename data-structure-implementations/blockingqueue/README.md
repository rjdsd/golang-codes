# Implementation Of Blocking Queue

* Enqueue operation will Wait if the Queue is Full. 
* Any Dequeue operation will Wait, if the Queue is Emtry. 
* During the wait, system will voluntarily release the locks to ensure Deadlock-Avoidance.

## Benchmarking

### System Spec
4 Core CPU, 8 GB RAM, Ubuntu 16.04

### Throuput

| Operation# | Time To Complete  |    Comment   
| -----------| ----------------  | ------------------------|
|  2000      | 0.027sec          | 1k Enque,  1k DQ        | 
|  20000     | 14.670sec         | 10k Enque, 10k DQ       |  
|  100,000   | 234.640s          | 50k Enque, 50k DQ       |        
