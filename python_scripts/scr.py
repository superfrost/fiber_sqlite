import sys
import os
import numpy as np
from PIL import Image

work_dir = './public/img/'
fileName = sys.argv[1]
sizeOfSegment = sys.argv[2]
insertedImageId = sys.argv[3]
flagCreateThumbnail = int(sys.argv[4])

dirForThumbnails = work_dir + "mini/"
if (flagCreateThumbnail):
  thumbImg = Image.open(work_dir + fileName)
  thumbImg.thumbnail((300, 300))
  thumbImg.save(dirForThumbnails + fileName, quality=70, subsampling=0)

try:
  sys.stdout.write(fileName)
except:
  sys.stderr.write("Some Error")
  # pass