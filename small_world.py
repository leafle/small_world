def small_world(points):
  ret = []
  for i in range(len(points)):
    cl3 = []
    for j in range(len(points)):
      if i == j: continue
      cl3.append((j, getDistanceSq(points[i], points[j])))
      cl3 = sorted(cl3, key=lambda p: p[1])
      if len(cl3) > 3: cl3.pop()
    ret.append([cl3[0][0], cl3[1][0], cl3[2][0]])
  return ret

distances = {}
def getDistanceSq(p1, p2):
  x = p1[0] - p2[0]
  y = p1[1] - p2[1]
  return x**2 + y**2


def small_world_half(points):
  ret = []
  for i in range(len(points)):
    cl3 = []
    for j in range(len(points)):
      if i == j: continue
      cl3.append((j, getDistanceSq(points[i], points[j])))
      cl3 = sorted(cl3, key=lambda p: p[1])
      if len(cl3) > 3: cl3.pop()
    ret.append([cl3[0][0], cl3[1][0], cl3[2][0]])
  return ret

print small_world([[1,2],[3,4],[2,3],[3,5]])
