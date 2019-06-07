import numpy as np
import pandas as pd
import matplotlib.pyplot as plt

dfInsert = pd.read_csv('mysql/10000_insert.csv')
dfRead = pd.read_csv('mysql/10000_read.csv')
dfInsertLonghorn = pd.read_csv('mysql/longhorn_10000_insert.csv')
dfReadLonghorn = pd.read_csv('mysql/longhorn_10000_read.csv')
dfInsertNFS = pd.read_csv('mysql/nfs_10000_insert.csv')
dfReadNFS = pd.read_csv('mysql/nfs_10000_read.csv')

# data to plot
n_groups = 2
class_a = [dfInsert.describe().iloc[1][0], dfRead.describe().iloc[1][0]]
class_b = [dfInsertLonghorn.describe().iloc[1][0], dfReadLonghorn.describe().iloc[1][0]]
class_c = [dfInsertNFS.describe().iloc[1][0], dfReadNFS.describe().iloc[1][0]]

# create plot
fig, ax = plt.subplots()
index = np.arange(n_groups)
bar_width = 0.2
opacity = 0.8

rects1 = plt.bar(index, class_a, bar_width,
alpha=opacity,
# color='b',
label='No persistence')

rects2 = plt.bar(index + bar_width, class_b, bar_width,
alpha=opacity,
# color='g',
label='Longhorn')

rects3 = plt.bar(index + 2*bar_width, class_c, bar_width,
alpha=opacity,
# color='y',
label='NFS')

plt.ylabel("time (sec)")
plt.title('MySQL Insert and Read Mean')
plt.xticks(index + bar_width, ('Insert', 'Read'))
plt.legend()

plt.tight_layout()
plt.show()


# data to plot
n_groups = 2
class_a = [dfInsert.describe().iloc[4][0], dfRead.describe().iloc[4][0]]
class_b = [dfInsertLonghorn.describe().iloc[4][0], dfReadLonghorn.describe().iloc[4][0]]
class_c = [dfInsertNFS.describe().iloc[4][0], dfReadNFS.describe().iloc[4][0]]

# create plot
fig, ax = plt.subplots()
index = np.arange(n_groups)
bar_width = 0.2
opacity = 0.8

rects1 = plt.bar(index, class_a, bar_width,
alpha=opacity,
# color='b',
label='No persistence')

rects2 = plt.bar(index + bar_width, class_b, bar_width,
alpha=opacity,
# color='g',
label='Longhorn')

rects3 = plt.bar(index + 2*bar_width, class_c, bar_width,
alpha=opacity,
# color='y',
label='NFS')

plt.ylabel("time (sec)")
plt.title('MySQL Insert and Read Q1')
plt.xticks(index + bar_width, ('Insert', 'Read'))
plt.legend()

plt.tight_layout()
plt.show()