import numpy as np
import pandas as pd
import matplotlib.pyplot as plt

small_upload_df_dict = dict()
small_upload_df_dict['File upload (No persistence)'] = pd.read_csv('minio/small_file_upload.csv')
small_upload_df_dict['File upload (Longhorn)'] = pd.read_csv('minio/small_file_longhorn_upload.csv')
small_upload_df_dict['File upload (NFS)'] = pd.read_csv('minio/small_file_nfs_upload.csv')

small_download_df_dict = dict()
small_download_df_dict['File download (No persistence)'] = pd.read_csv('minio/small_file_download.csv')
small_download_df_dict['File download (Longhorn)'] = pd.read_csv('minio/small_file_longhorn_download.csv')
small_download_df_dict['File download (NFS)'] = pd.read_csv('minio/small_file_nfs_download.csv')

big_upload_df_dict = dict()
big_upload_df_dict['File upload (No persistence)'] = pd.read_csv('minio/big_file_upload.csv')
big_upload_df_dict['File upload (Longhorn)'] = pd.read_csv('minio/big_file_longhorn_upload.csv')
big_upload_df_dict['File upload (NFS)'] = pd.read_csv('minio/big_file_nfs_upload.csv')

big_download_df_dict = dict()
big_download_df_dict['File download (No persistence)'] = pd.read_csv('minio/big_file_download.csv')
big_download_df_dict['File download (Longhorn)'] = pd.read_csv('minio/big_file_longhorn_download.csv')
big_download_df_dict['File download (NFS)'] = pd.read_csv('minio/big_file_nfs_download.csv')

color = ['g', 'b', 'r']

for idx, dic in enumerate([small_upload_df_dict, small_download_df_dict, big_upload_df_dict, big_download_df_dict]):
    for key in dic.keys():
        fig, ax = plt.subplots()
        mean = dic[key].mean()
        mean.index = np.arange(1,len(mean)+1)
        mean.plot(ax=ax, style='o-')
        dic[key].boxplot(showfliers=True, ax=ax)
        ax.set_title(key)
        if idx > 1:
            ax.set_xlabel('files (64MB per file)')
        else:
            ax.set_xlabel('file size (MB)')
        ax.set_ylabel('time (sec)')
        plt.show()

    fig, ax = plt.subplots()
    for c, key in enumerate(dic.keys()):
        mean = dic[key].mean()
        if idx > 1:
            mean.index = np.arange(1*64, (len(mean)+1)*64, 64)
        else:
            mean.index = np.arange(1, len(mean)+1)
        
        mean.plot(ax=ax, style=color[c]+'o-', label=key + ' Mean')
        q1 = dic[key].quantile(q=0.25, axis=0, numeric_only=True, interpolation='linear')
        if idx > 1:
            q1.index = np.arange(1*64, (len(q1)+1)*64, 64)
        else:
            q1.index = np.arange(1, len(q1)+1)
        q1.plot(ax=ax, style=color[c]+'x--', label=key + ' Top 25%')
        if idx > 1:
            ax.set_xlabel('files (64MB per file)')
        else:
            ax.set_xlabel('file size (MB)')
        ax.set_ylabel('time (sec)')
    plt.legend(loc='upper left')
    plt.show()
