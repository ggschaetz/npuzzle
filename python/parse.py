#!/usr/bin/env python

# **************************************************************************** #
#                                                                              #
#                                                         :::      ::::::::    #
#    main.py                                            :+:      :+:    :+:    #
#                                                     +:+ +:+         +:+      #
#    By: gschaetz <gschaetz@student.42.fr>          +#+  +:+       +#+         #
#                                                 +#+#+#+#+#+   +#+            #
#    Created: 2018/10/09 14:14:46 by gschaetz          #+#    #+#              #
#    Updated: 2018/10/09 14:14:49 by gschaetz         ###   ########.fr        #
#                                                                              #
# **************************************************************************** #

import sys
import os
import numpy as np

class Parser():

    def __init__(self, args):
        self.file_name = args.file_name
        self.heuristique = args.heur
        self.size = 0
        self.map_in_list = list()
        self.read_file()

    def read_file(self):
        try :
            self.file = open(self.file_name, mode='r')
            if os.stat(self.file_name).st_size == 0:
                print ("\"{file_name}\" is empty".format(file_name=self.file_name))
                sys.exit(1)
        
        except IOError :
            print ("can't open \"{file_name}\"".format(file_name=self.file_name))
            sys.exit(1)
        
    def parse_input(self):
        """parse imput map"""

        def error(name):
            switcher = {
                'matrice': 'bad matrice format',
                'size': 'bad size format'
            }
            print("{}".format(switcher[name]))
            sys.exit(0)

        def drop_white(lst):
            """drop empty boxe of lst"""
            i = 0
            while i < len(lst):
                if lst[i] == '':
                    del lst[i]
                else:    
                    i += 1
            return lst

        def int_list(lst):
            """convert list(,str()) in list(,int())"""
            i = 0
            l = len(lst)
            while i < l:
                lst[i] = int(lst[i])
                i += 1
            return lst 
        
        def list_isnumeric(lst):
            """check if all members of list is numeric"""
            for member in lst:
                if not member.isnumeric():
                    return False
            return True

        def check_size_map(size):
            """check if the line containing the size is formatted 
            correctly and return it then""" 
            i = 0
            size = size.split(' ')
            while i < len(size):
                if size[i] == '':
                    size.pop(i)
                else:
                    i += 1
            if len(size) != 1 or int(size[0]) < 3:
                error('size')
            return int(size[0])

        def format_line(lst):
            """ format line matrice to a list """
            line = lst[0].split(' ')
            line = drop_white(line)
            if list_isnumeric(line) == True and len(line) == self.size:
                return line
            error('matrice')

        i = 0
        for line in self.file:
            lst = line.split('#')
            if lst[0] != '':
                i += 1
            else:
                continue
            lst[0] = lst[0].replace("\n", "")
            if i == 1:
                self.size = check_size_map(lst[0])
            else:
                self.map_in_list.append(int_list(format_line(lst)))