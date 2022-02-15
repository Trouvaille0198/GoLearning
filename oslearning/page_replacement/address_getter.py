import os
import re
import time
import string
import psutil

# 统计某一个进程名所占用的内存，同一个进程名，可能有多个进程


def countProcessMemoey(processName):
    pattern = re.compile(r'([^\s]+)\s+(\d+)\s.*\s([^\s]+\sK)')
    cmd = 'tasklist /fi "imagename eq ' + processName + \
        '"' + ' | findstr.exe ' + '"' + processName + '"'
    # findstr后面的程序名加上引号防止出现程序带有空格
    result = os.popen(cmd).read()
    resultList = result.split("\n")
    pids = []
    for srcLine in resultList:
        srcLine = "".join(srcLine.split('\n'))
        if len(srcLine) == 0:
            break
        m = pattern.search(srcLine)
        if m == None:
            continue
        # 由于是查看python进程所占内存，因此通过pid将本程序过滤掉
        if str(os.getpid()) == m.group(2):
            continue
        ori_mem = m.group(3).replace(',', '')
        ori_mem = ori_mem.replace(' K', '')
        ori_mem = ori_mem.replace(r'\sK', '')
        memEach = int(ori_mem)
        # print ('ProcessName:'+ m.group(1) + '\tPID:' + m.group(2) + '\tmemory size:%.2f'% (memEach * 1.0 /1024), 'M')
        # print(time.ctime())  #打印当前的时间
        # print('PID:' + m.group(2))
        pids.append(m.group(2))
    return pids
    # print("*" * 58)


if __name__ == '__main__':
    lis = []
    pids = psutil.pids()
    for pid in pids:
        p = psutil.Process(pid)
        lis.append(str(p.name()))
    results = []
    for i in lis[:50]:
        ProcessName = i  # 进程名
        results.extend(countProcessMemoey(ProcessName))
    print(results)
    with open("./pids.txt", "w") as f_b:
        f_b.write(",".join(results))
    # ProcessName = 'TabNine.exe'
    # countProcessMemoey(ProcessName)
