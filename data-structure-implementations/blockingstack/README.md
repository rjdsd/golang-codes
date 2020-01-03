# Implementation Of Blocking Stack

* Push operation will Wait if the Stack is Full. 
* Any Pop operation will Wait if the Stack is Emtry. 
* During the wait, system will voluntarily release the locks to ensure Deadlock-Avoidance.

## Benchmarking

### System Spec
4 Core CPU, 8 GB RAM, Ubuntu 16.04

### Throuput

| Operation# | Time To Complete  |    Comment   
| -----------| ----------------  | ------------------------|
|  2000      | 0.003  sec        | 1k Push,  1k Pop        | 
|  20000     | 17.293 sec        | 10k Push, 10k Pop       |     
