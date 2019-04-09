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

import sys
import math

def find_val_in_map(val0, state, map):
    """ find i, j in map for one value (val0) in input """
    #print ("state la = ", state)
    #print ("type state la = ", type(state))
    #print ("size map = ", map.size)
    i = 0
    while i < map.size:
        j = 0
        while j < map.size:
            #print ("val0 = ", val0, " state[",i,"][",j,"] = " ,state[i][j])
            if val0 == state[i][j]:
                #print ("return")
                return i, j
            j += 1
        i += 1
    sys.exit(1)

def sum_heuristique(map, state):
    """ comput sum of heurstique for state """
    if map.heuristique == 'Manhatan':
        def heuristique(xa, xb, ya, yb):
            return abs(xb - xa) + abs(yb - ya)
    elif map.heuristique == "Euclidiene":
        def heuristique(xa, xb, ya, yb):
            return math.sqrt((xa - xb) * (xa - xb) + ((ya - yb) * (ya - yb)))
    elif map.heuristique == "Greedy":
        def heuristique(xa, xb, ya, yb):
            if abs(xb - xa) + abs(yb - ya) != 0:
                return 1
            return 0
    elif map.heuristique == "Missplace":
        state_in_list = list()
        final_in_list = list()
        miss_place = 0
        i = 0
        state_in_list = map.build_list(state, state_in_list)
        final_in_list = map.build_list(map.final_map, final_in_list)
        for term in state_in_list:
            if term != final_in_list[i]:
                miss_place += 1
            i += 1
        return miss_place
    
    if map.heuristique == "Inversion":
        state_in_list = list()
        state_in_list = map.build_list(state, state_in_list)
        state_in_list = [term for term in state_in_list if term != 0]
        nb_inversions = map.compute_inversion(state_in_list)
        return (nb_inversions)

    sum_heuristique_final = 0
    sum_heuristique_start = 0
    xa = 0
    while xa < map.size:
        ya = 0
        while ya < map.size:
            xb, yb = find_val_in_map(map.final_map[xa][ya], state, map)
            sum_heuristique_final += heuristique(xa, xb, ya, yb) ##math.sqrt((xa - xb) * (xa - xb) + ((ya - yb) * (ya - yb))) ##abs(xb - xa) + abs(yb - ya)
            xb, yb = find_val_in_map(map.start_map[xa][ya], state, map)
            sum_heuristique_start += heuristique(xa, xb, ya, yb) ##math.sqrt((xa - xb) * (xa - xb) + ((ya - yb) * (ya - yb))) ##abs(xb - xa) + abs(yb - ya)
            ya += 1
        xa += 1

    return sum_heuristique_final
            
