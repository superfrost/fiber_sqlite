import os
import sys
import numpy as np
from PIL import Image

# Command example:
# python .\segmentation.py image.jpg 8 1 1
# Output: 0 (if all ok) or 1 (if error)
# if ok then 'result-1-8.jpg' will be created in work directory and 
# if set flag generate thumbnail then image.jpg will be created in work_dir + "mini/" directory

# Arguments from commandline
work_dir = './public/img/'
dirForResult = work_dir + "result/"
dirForThumbnails = work_dir + "mini/"
dirForResultThumbnails = dirForThumbnails + "result/"

fileName = sys.argv[1]
sizeOfSegment = sys.argv[2]
insertedImageId = sys.argv[3]
flagCreateThumbnail = int(sys.argv[4])

# Create directories for thumbnails and result
try:
  os.mkdir(work_dir + "mini/result/")
  os.mkdir(work_dir + "result/")
except:
  pass


# Create thumbnail for image

if (flagCreateThumbnail):
  thumbImg = Image.open(work_dir + fileName)
  thumbImg.thumbnail((300, 300))
  thumbImg.save(dirForThumbnails + fileName, quality=70, subsampling=0)

# Segmentation script
img = Image.open(work_dir + fileName)
gray_img = img.convert("L")
np_array = np.array(gray_img, dtype='uint8')

def segmentation(npArray, segmentSize = 8):
  
  """Convert image in numpy.array format to binary image numpy.array format where min value = 0 and max value = 255

  Returns:
      [type]: [description]
  """
  widthArr = npArray.shape[0]
  heightArr = npArray.shape[1]
  
  widthSteps = widthArr // segmentSize
  heightSteps = heightArr // segmentSize

  newArr = np.zeros((widthSteps * segmentSize, heightSteps * segmentSize))
  
  for i in range(widthSteps):
    for j in range(heightSteps):
      segment = npArray[i*segmentSize : i*segmentSize+segmentSize, j*segmentSize : j*segmentSize+segmentSize]
      meanOfSegment = np.mean(segment)
      normSegment = segment - meanOfSegment
      poweredSegment = normSegment ** 2
      filteredSegment = (poweredSegment > 0) * poweredSegment
      newArr[i*segmentSize : i*segmentSize+segmentSize, j*segmentSize : j*segmentSize+segmentSize] = filteredSegment

  newArr = newArr - np.mean(newArr)
  newArr = (newArr > 0) * 255
  # newArr = newArr // np.mean(newArr)
  # newArr = (newArr > 0) * 255

  return newArr

try:
  segmentedImg = segmentation(np_array, int(sizeOfSegment))

  img_from_npArray = Image.fromarray(segmentedImg.astype('uint8'), 'L')

  resultFileName = 'result-' + insertedImageId + '-' + sizeOfSegment + '.jpg'
  save_dir = dirForResult + resultFileName
  img_from_npArray.save(save_dir, quality=95, subsampling=0)

  # Thumbnail for result image after segmentation
  if (flagCreateThumbnail):
    img_from_npArray.thumbnail((300,300))
    img_from_npArray.save(dirForResultThumbnails + resultFileName, quality=70, subsampling=0)
  
  sys.stdout.write(resultFileName)
except:
  sys.stderr.write("Some error")
  sys.exit(1)
  # pass





