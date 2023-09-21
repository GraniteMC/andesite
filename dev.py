import argparse, shutil, os
from watchdog.observers import Observer
from watchdog.events import FileSystemEventHandler

parser = argparse.ArgumentParser(description='Some scripts for packed')
parser.add_argument('--scss', '-c', '--css', action='store_true', help='compile scss to css')
parser.add_argument('--js', '-c', '--css', action='store_true', help='compile ts to js')
parser.add_argument('--csswatch', '-w', action='store_true', help='watch css')
parser.add_argument('--ts', '-t', action='store_true', help='watch ts')

args = parser.parse_args()
if args.scss:

    cmds = [
        'sass static/index.scss static/index.css'
    ]

    for cmd in cmds:
        print(f'Compiling {cmd.split(" ")[1]}')
        os.system(cmd)

if args.js:

    cmds = [
        'tsc static/index.ts'
    ]

    for cmd in cmds:
        print(f'Compiling {cmd.split(" ")[1]}')
        os.system(cmd)

if args.csswatch:
    
    cmds = [
        'sass --watch static/index.scss static/index.css'
    ]
    
    for cmd in cmds:
        print(f'Watching {cmd.split(" ")[2]}')
        os.system(cmd)

if args.ts:

    #get all ts files in public/ts
    ts_files = [f for f in os.listdir('static') if f.endswith('.ts')]

    #listen for chasnges in all ts files
    for ts_file in ts_files:
        os.system(f'tsc -w static/{ts_file} --outDir static')
