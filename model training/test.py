import praw
from nltk.stem import PorterStemmer
from nltk.corpus import stopwords
import pandas as pd
import re
from sklearn.model_selection import train_test_split
from sklearn.feature_extraction.text import TfidfVectorizer
from sklearn.ensemble import RandomForestClassifier
from sklearn.pipeline import Pipeline
from sklearn.feature_selection import SelectKBest, chi2
from sklearn import preprocessing
from sklearn.svm import LinearSVC
from sklearn.neighbors import KNeighborsClassifier
import numpy as np