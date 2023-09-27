# Reversing a String

## Problem Description

Please implement an algorithm to reverse a given string without using **additional data structures and storage space** (a single process variable can be used).

Given a string, return a string that is the reversed version of the original. Ensure that the length of the string is less than or equal to 5000.

## Solution Approach

Reversing a string essentially means flipping it around its central character. That is, assign `str[len-1]` to `str[0]`, and assign `str[0]` to `str[len-1]`.
