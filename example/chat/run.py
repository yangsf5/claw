#!/usr/bin/python
# Author: sheppard(ysf1026@gmail.com) 2014-03-21

import getopt
import os
import sys

def usage():
    print 'usage:'
    print '-h\n\thelp'
    print '-t server type\n\t(master or harbor1 or harbor2)'

def main(argv):
    try:
        opts, args = getopt.getopt(argv, 'ht:')
    except getopt.GetoptError:
        usage()
        sys.exit(2)

    for opt, arg in opts:
        if opt == '-h':
            usage()
            sys.exit()
        elif opt == '-t':
            global _type
            _type = arg

    start()

def start():
    try:
        _type
    except:
        usage()
        sys.exit(2)

    if _type == 'master':
        os.system('go run chat.go -logtostderr=true -config="./config/config_master.xml"')
    elif _type == 'harbor1':
        os.system('go run chat.go -logtostderr=true -config=./config/config_harbor1.xml')
    elif _type == 'harbor2':
        os.system('go run chat.go -logtostderr=true -config=./config/config_harbor2.xml')
    else:
        usage()

if __name__ == '__main__':
    main(sys.argv[1:])
