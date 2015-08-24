#!/usr/bin/python

import getopt
import os
import sys

def usage():
    print 'usage:'
    print '-h\n\thelp'
    print '-t server type\n\t(mock_server or master or harbor1 or harbor2)'

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

    common_str = 'go run stress/stress.go -logtostderr=true -config=./stress/config/';

    if _type == 'mock_server':
        os.system('go run mock_server/mock_server.go')
    elif _type == 'master':
        os.system(common_str + 'config_master.xml')
    elif _type == 'harbor1':
        os.system(common_str + '/config_harbor1.xml')
    elif _type == 'harbor2':
        os.system(common_str + '/config_harbor2.xml')
    else:
        usage()

if __name__ == '__main__':
    main(sys.argv[1:])
