---
layout: default
title: Vars
nav_order: 2
parent: Feature
---

<link rel="stylesheet" href="../../../assets/css/custom.css">

# Vars

By making use of block `vars`,  we define variables that could be reused by one or more scenarios. This block contains a 
set of attributes in hcl format. Only one block `vars` is permitted per file. 

**Example** [download](https://raw.githubusercontent.com/wesovilabs/orion/master/examples/vars/feature001.hcl)

```hcl
vars {
  x = 1
  elements = [
    { product = "tofu",     vegan  = true},
    { product = "meat",     vegan  = false},
    { product = "fish",     vegan  = false},
    { product = "avocado",  vegan  = true},
  ]
}
scenario "feature with plain var" {
  when "multiply by 2" {
    set result{
      value = x * 2
    }
  }
  then "check output" {
    assert {
      assertion = result==2
    }
  }
}
scenario "feature with array var" {
  given "initial karma"{
    set karma {
      value = 0
    }
  }
  when "calculate karma" {
    block {
      set karma {
        value = karma + 1
        when = elements[_.index].vegan
      }
      set karma {
        value = karma - 1
        when = !elements[_.index].vegan
      }
      count = len(elements)
    }
  }
  then "check output" {
    assert {
      assertion = karma==0
    }
  }
}

``` 
