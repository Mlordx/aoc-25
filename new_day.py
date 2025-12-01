#!/usr/bin/python3
import subprocess
import sys

day = sys.argv[1]
lang = sys.argv[2] if len(sys.argv) > 1 else "py"
day_with_two_digits = '0' + day if int(day) < 10 else day
cmd = 'mkdir day{}'.format(day_with_two_digits)
subprocess.check_output(cmd, shell=True)
cmd = 'touch day{}/day{}.{}'.format(day_with_two_digits, day_with_two_digits,  lang)
subprocess.check_output(cmd, shell=True)
cmd = './get_input.py --day {} > day{}/input.txt'.format(day, day_with_two_digits)
output = subprocess.check_output(cmd, shell=True)
output = output.decode('utf-8')
print(output, end='')
print('\n'.join(output.split('\n')[:10]), file=sys.stderr)