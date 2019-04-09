#!/usr/bin/env python

# **************************************************************************** #
#                                                                              #
#                                                         :::      ::::::::    #
#    map                                                :+:      :+:    :+:    #
#                                                     +:+ +:+         +:+      #
#    By: gschaetz <gschaetz@student.42.fr>          +#+  +:+       +#+         #
#                                                 +#+#+#+#+#+   +#+            #
#    Created: 2018/10/09 16:37:15 by gschaetz          #+#    #+#              #
#    Updated: 2018/10/09 16:37:15 by gschaetz         ###   ########.fr        #
#                                                                              #
# **************************************************************************** #

import numpy as np

from parse import Parser

class Map(Parser):

    def __init__(self, args):
        Parser.__init__(self, args)
        self.heuristique = args.heur
        self.size = self.size
        self.solvable = True
        self.start_map = np.zeros(self.size)
        self.final_map = np.zeros(self.size)
        self.curent_state = (np.zeros(self.size), self.size * self.size)

    def solvable(self):
        print ("solvable")
    
    def comput_final_map(self):
        """generate final map/node for A*"""    
        def filing_right(x, y, i, nb_case):
            while y < self.size and self.final_map[x][y] == 0 and i != nb_case:
                self.final_map[x][y] = i
                i += 1
                y += 1
            y -= 1
            x += 1
            return x, y, i
    
        def filing_down(x, y, i, nb_case):
            while x < self.size and self.final_map[x][y] == 0 and i != nb_case:
                self.final_map[x][y] = i
                i += 1
                x += 1
            x -= 1
            y -= 1
            return x, y, i
        
        def filing_left(x, y, i, nb_case):
            while y >= 0 and self.final_map[x][y] == 0 and i != nb_case:
                self.final_map[x][y] = i
                i += 1
                y -= 1
            y += 1
            x -= 1
            return x, y, i

        def filing_up(x, y, i, nb_case):
            while x >= 0 and self.final_map[x][y] == 0 and i != nb_case:
                self.final_map[x][y] = i
                i += 1
                x -= 1
            x += 1
            y += 1
            return x, y, i
        
        nb_case = self.size * self.size
        i = 1
        y = 0
        x = 0
        self.final_map = np.zeros((self.size, self.size))
        while i != nb_case:
            x, y, i = filing_right(x, y, i, nb_case)
            x, y, i = filing_down(x, y, i, nb_case)
            x, y, i = filing_left(x, y, i, nb_case)
            x, y, i = filing_up(x, y, i, nb_case)

    def read_left_right(self, x, y, i, recept_list, map):
        while y < self.size - i:
            recept_list.append(map[x][y])
            y += 1
        y -= 1
        x += 1
        return x, y, recept_list

    def read_top_bot(self, x, y, i, recept_list, map):
        while x < self.size - i:
            recept_list.append(map[x][y])
            x += 1
        x -= 1
        y -= 1
        return x, y, recept_list

    def read_right_left(self, x, y, i, recept_list, map):
        while y >= 0 + i:
            recept_list.append(map[x][y])
            y -= 1
        y += 1
        x -= 1
        return x, y, recept_list

    def read_bot_top(self, x, y, i, recept_list, map):
        while x > 0 + i:
            recept_list.append(map[x][y])
            x -= 1
        x += 1
        y += 1
        return x, y, recept_list


    def build_list(self, map, recept_list):
        """ convert map array to list:
            [[1,2,3],
                [8,0,4],
                [7,6,5]]
            
            egual to [1,2,3,4,5,6,7,8,0] """
        #print("entre = ", type(recept_list), print(recept_list))
        i = 0
        x = 0
        y = 0
        while (i < (self.size + 1)//2):
            x, y, recept_list = self.read_left_right(x, y, i, recept_list, map)
            x, y, recept_list = self.read_top_bot(x, y, i, recept_list, map)
            x, y, recept_list = self.read_right_left(x, y, i, recept_list, map)
            x, y, recept_list = self.read_bot_top(x, y, i, recept_list, map)
            i += 1
        return recept_list

    def find_index_of_term_in_list(self, term, lst):
        i = 0
        for member in lst:
            if member == term:
                return i
            i += 1
        return -1
    
    def compute_inversion(self, start_list):
        i = 0
        len_list = len(start_list)
        nb_inversions = 0
        while (i < len_list):
            j = i
            while (j < len_list):
                if start_list[i] > start_list[j]:
                    nb_inversions += 1
                j += 1
            i += 1
        return nb_inversions

    def odd_size(self, start_list):
        if compute_inversion(start_list) % 2 == 0:
            return True
        return False

    def even_size(self, start_list, zero_position):
        if zero_position % 2 == 0 and compute_inversion(start_list) % 2 != 0:
            return True
        elif zero_position % 2 != 0 and compute_inversion(start_list) % 2 == 0:
            return True
        return False 

    def check_solvable(self):
        """ check if the imput map is sovable or not """

        

        start_list = list()
        final_list = list()
        start_list = self.build_list(self.start_map, start_list)
        final_list = self.build_list(self.final_map, final_list)
        zero_position_start = self.size - (self.find_index_of_term_in_list(0, start_list) // self.size)
        zero_position_final = self.size - (self.find_index_of_term_in_list(0, final_list) // self.size)
        start_list = [term for term in start_list if term != 0]
        final_list = [term for term in final_list if term != 0]
        inversions_start = self.compute_inversion(start_list)
        inversions_final = self.compute_inversion(final_list)
        if (inversions_start % 2 != inversions_final % 2):
            print ("map not solvable")
            exit(0)

    def compute_start_map(self):
        self.start_map = np.array(self.map_in_list)