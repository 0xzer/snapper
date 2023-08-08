import subprocess
import os


def find_proto_files(src):
    paths = []
    for root, dirs, files in os.walk(src):
        for file in files:
            if file.endswith(".proto"):
                full_path = os.path.join(root, file)
                paths.append(full_path)
    return paths



src = os.path.join(os.getcwd(), 'files')
paths = find_proto_files(src)

for path in paths:
    cmd = "protoc --proto_path=/home/zero/Desktop/snapper/protos/files --go_out=. " + path
    subprocess.run(cmd, shell=True)