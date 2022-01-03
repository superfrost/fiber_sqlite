import sys
import os
# import numpy as np
# from PIL import Image

cwd = os.getcwd()

work_dir = './public/img/'
fileName = sys.argv[1]
sizeOfSegment = sys.argv[2]
insertedImageId = sys.argv[3]
flagCreateThumbnail = int(sys.argv[4])

outString = cwd + " " + work_dir #+ " " + fileName + " " + sizeOfSegment + " " + insertedImageId + " " + flagCreateThumbnail
try:
  sys.stdout.write(outString)
except:
  sys.stderr.write("Some Error")
  # pass