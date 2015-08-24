#!/usr/bin/python

import getopt
import os
import sys

def usage():
    print 'usage:'
    print '-h\n\thelp'
    print '-t server type\n\t(server or client)'

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

    common_str = ' -logtostderr=true -config=./config/config_master.xml';
    if _type == 'server':
        os.system('go run ./server/server.go' + common_str)
    elif _type == 'client':
        os.system('go run client.go' + common_str)
    else:
        usage()

if __name__ == '__main__':
    main(sys.argv[1:])
