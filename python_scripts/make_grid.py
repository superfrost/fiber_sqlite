import sys
import math
from PIL import Image, ImageDraw, ImageFont

# Command example:
# python .\make_grid.py "./images/" Item_1.jpg Item_2.jpg Item_3.jpg Item_3.jpg Item_3.jpg 1 2 3 4 5
# Output: 0 (if all ok) or 1 (if error)

work_dir = sys.argv[1]

args = sys.argv[2:]
argsLen = len(args)
pivot = (argsLen // 2) + 2

fileNames = sys.argv[2:pivot]
filesId = sys.argv[pivot:]

def createGrid(workdir, arrayFileNames, arrayFileId = []):
  try:
    thumb_size = 300
    gap = 20
    images_count = len(arrayFileNames)
    x_grid_size = math.ceil(math.sqrt(images_count))
    y_grid_size = round(math.sqrt(images_count))
    x_grid_size_pixel = thumb_size * x_grid_size
    y_grid_size_pixel = thumb_size * y_grid_size

    grid_image = Image.new('RGB', (x_grid_size_pixel + gap * (x_grid_size - 1), y_grid_size_pixel + gap * (y_grid_size)), (255, 255, 255))

    index = 0
    for file in arrayFileNames:
      x_position = index % x_grid_size
      y_position = index // x_grid_size
      thumbImg = Image.open(workdir + file)
      thumbImg.thumbnail((thumb_size, thumb_size))
      grid_image.paste(thumbImg, (x_position * thumb_size + gap * x_position, y_position * thumb_size + gap * y_position))

      index += 1
    
    drawImg = ImageDraw.Draw(grid_image)
    imageFont = ImageFont.truetype('./fonts/OpenSans-Regular.ttf', 20)

    index = 0
    for id in arrayFileId:
      x_position = index % x_grid_size
      y_position = index // x_grid_size

      text_x_position = x_position * thumb_size + thumb_size // 2 + gap * x_position - 50
      text_y_position = y_position * thumb_size + thumb_size + gap * y_position - 5
      drawImg.text((text_x_position, text_y_position), "Id: " + id, fill=(0, 0, 0), font=imageFont)

      index += 1

    grid_image.save(work_dir + 'grid/' + 'grid.jpg', quality=70, subsampling=0)
    
    # sys.stdout.write('0')
    sys.exit(0)
  except:
    # sys.stdout.write('1')
    sys.exit(1)

createGrid(work_dir, fileNames, filesId)