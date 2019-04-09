# **************************************************************************** #
#                                                                              #
#                                                         :::      ::::::::    #
#    a_star_v1.py                                       :+:      :+:    :+:    #
#                                                     +:+ +:+         +:+      #
#    By: gschaetz <gschaetz@student.42.fr>          +#+  +:+       +#+         #
#                                                 +#+#+#+#+#+   +#+            #
#    Created: 2019/02/05 16:03:31 by gschaetz          #+#    #+#              #
#    Updated: 2019/02/27 11:16:30 by gschaetz         ###   ########.fr        #
#                                                                              #
# **************************************************************************** #

import numpy as np
import time
import sys
import collections

import heuristique as hrst

def find_blanc(map):
    i = 0
    for line in map:
        j = 0
        for val in line:
            if val == 0:
                return (i, j)
            j += 1
        i += 1

def map_swap_bot(map, blanc):
    i, j = blanc[0], blanc[1] 
    new_map = np.copy(map)
    new_map[i][j] = map[i + 1][j]
    new_map[i + 1][j] = 0
    return new_map
    
def map_swap_top(map, blanc):
    i, j = blanc[0], blanc[1] 
    new_map = np.copy(map)
    new_map[i][j] = map[i - 1][j]
    new_map[i - 1][j] = 0
    return new_map

def map_swap_left(map, blanc):
    i, j = blanc[0], blanc[1] 
    new_map = np.copy(map)
    new_map[i][j] = map[i][j - 1]
    new_map[i][j - 1] = 0
    return new_map

def map_swap_right(map, blanc):
    i, j = blanc[0], blanc[1] 
    new_map = np.copy(map)
    new_map[i][j] = map[i][j + 1]
    new_map[i][j + 1] = 0
    return new_map

def next_state_node(map, state_node, cout):
    blanc = find_blanc(state_node)
    next_states = list()
    if blanc[0] != 0:
        state = map_swap_top(state_node, blanc)
        heuris = hrst.sum_heuristique(map, state)
        next_states.append((heuris, cout, heuris - cout, state, state_node))    
    if blanc[0] != len(state_node) - 1:
        state = map_swap_bot(state_node, blanc)
        heuris = hrst.sum_heuristique(map, state)
        next_states.append((heuris, cout, heuris - cout, state,  state_node))
    if blanc[1] != 0:
        state = map_swap_left(state_node, blanc)
        heuris = hrst.sum_heuristique(map, state)
        next_states.append((heuris, cout, heuris - cout, state, state_node))
    if blanc[1] != len(state_node) - 1:
        state = map_swap_right(state_node, blanc)
        heuris = hrst.sum_heuristique(map, state)
        next_states.append((heuris, cout, heuris - cout, state, state_node))
    
    ####### print states #######
    #print_states(state_node, next_states)
    ####### end print #######

    return next_states

def compare_node(hrst_node0, hrst_node1):
    if hrst_node0 > hrst_node1:
        return 1
    elif hrst_node0 == hrst_node1:
        return 0
    else:
        return -1

def add_close_list(map, state, open_list, close_list):
    close_list[str(state[0])] = (hrst.sum_heuristique(map, state[0]), str(state[0]))
    if not open_list.pop(str(state[0]), False):
        print("error, state not found in open list")
        sys.exit(1)
    return close_list

def add_adjacent(map, state, close_list, open_list):
    adjacent = next_state_node(map, state[3], state[1])
    for (heuristique, cout, real_hrst, node, previous_node) in adjacent:
        if (str(node) in close_list and cout >= state[1]) or (str(node) in open_list and cout >= state[1]):
            pass
        else:
            cout += 1
            heuristique += cout
            real_hrst = heuristique - cout
            open_list[str(node)] = (heuristique, cout, real_hrst, node, previous_node)

def str_to_list(string):
    lst = list()
    lst = string[1:-1].split("\n")
    new = list()
    for member in lst:
        ss_lst = list()
        member = member.split(" ")
        for val in member:
            val = val.strip('[')
            val = val.strip(']')
            if val != '':
                ss_lst.append(int(val))
        new.append(ss_lst)
    return new

def best_nodev2(open_list):
    for state, (heuristique, parent) in open_list.items():
        best_heuristique = (np.array(state), heuristique)
        break
    for state, (heuristique, parent) in open_list.items():
        if heuristique <= best_heuristique[1]:
            state = str_to_list(state)
            best_heuristique = (np.array(state), heuristique)
    return (best_heuristique)


def best_node(open_list, heuristique_state):
    best_heuristique = (np.zeros(0), heuristique_state)
    for state, (heuristique, parent) in open_list.items():
        if heuristique <= best_heuristique[1]:
            state = str_to_list(state)
            best_heuristique = (np.array(state), heuristique)
    return (best_heuristique)

def build_way(close_list):
    i = 0
    j = len(close_list) - 1 
    k = j
    items = list(close_list.items())
    rez = list()
    while (j > 0):
        rez.append(items[j][0])
        while (items[k][0] != str(items[j][1][4])):
            k -= 1
        while (j != k):
            j -= 1
    return (rez)
    
def find_first_key(open_list):
    i = 0
    for key, val in open_list.items():
        return key

def final_print(map, final_way, open_list, close_list, len_max_open_list, nb_node_checked):
    print("\n")
    print ("Maximum number of states ever represented in memory at the same time:             ", len_max_open_list)
    print ("Total number of states ever selected in the opened set:                           ", nb_node_checked)
    print ("Number of moves required to transition from the initial state to the final state: ", len(final_way))
    way = input("Do you want show resolution ?[y/n]")

    len_way = len(final_way) - 1 
    if way == 'y' or way == "yes":
        i = len_way
        while (i >= 0):
            print ("\nstate {state} : \n\n".format(state= 1 + (i - len_way) * - 1), final_way[i])
            i -= 1

def compute_priority_in_open_list(open_list):
    return collections.OrderedDict(sorted(open_list.items(), key=lambda x: x[1][2]))

def a_star(map):
    len_max_open_list = 0
    nb_node_checked = 0
    func_len_max_open_list = lambda open_list : len(open_list) if len(open_list) > len_max_open_list else len_max_open_list
    open_list = collections.OrderedDict()
    close_list = collections.OrderedDict()
    map.curent_state = (np.zeros(map.size), map.size * map.size)
    heuristique_start = hrst.sum_heuristique(map, map.start_map)
    open_list[str(map.start_map)] = (heuristique_start, 0, 0, map.start_map, None)
    map.curent_state = open_list[str(map.start_map)]
    while map.curent_state[0] != 0 or open_list != {}:
        len_max_open_list = func_len_max_open_list(open_list)  
        node = open_list.pop(find_first_key(open_list))
        nb_node_checked += 1
        if hrst.sum_heuristique(map, node[3]) == 0:
            close_list[str(node[3])] = (node[0], node[1], node[2], node[3], node[4])
            final_way = build_way(close_list)
            final_print(map, final_way, open_list, close_list, len_max_open_list, nb_node_checked)
            return
        add_adjacent(map, node, close_list, open_list)
        open_list = compute_priority_in_open_list(open_list)
        close_list[str(node[3])] = (node[0], node[1], node[2], node[3], node[4])