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

import map
import numpy as np

import heuristique as hrst
import a_star_v1 as star

def solve(map):
    """ pile = {map:heuristique} """
    pile = dict()
    map.numpy_map()
    pile[str(map.start_map)] = hrst.sum_heuristique(map, map.start_map)
    states = star.next_state_node(map.start_map)
    for matrice in states:
        pile[str(matrice.tolist())] = hrst.sum_heuristique(map, matrice)
    print (pile)

