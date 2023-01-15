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


reddit_read_only = praw.Reddit(client_id="xnsIxViGuyQrTKs7VnnYCQ",  # your client id
                               client_secret="LuRt_v69n1GxxADiaLssD25fWOkuyw",  # your client secret
                               user_agent="safespace",
                               redirect_uri="http://localhost:8080")

# Curated list of helpful subreddits
subredditList = ["TRAUMATOOLBOX",
                 "rapecounseling",
                 "7CupsofTea",
                 "addiction",
                 "ADHD",
                 "Advice",
                 "affirmations",
                 "afterthesilence",
                 "Agoraphobia",
                 "alcoholism",
                 "Anger",
                 "Antipsychiatry",
                 "Anxiety",
                 "ARFID",
                 "AskDocs",
                 "BipolarReddit",
                 "BodyAcceptance",
                 "bulimia",
                 "CompulsiveSkinPicking",
                 "dbtselfhelp",
                 "depression_help",
                 "disability",
                 "domesticviolence",
                 "EatingDisorders",
                 "EOOD",
                 "ForeverAlone",
                 "getting_over_it",
                 "HealthAnxiety",
                 "helpmecope",
                 "mentalhealth",
                 "mentalillness",
                 "Needafriend",
                 "OCD",
                 "offmychest",
                 "Psychiatry",
                 "ptsd",
                 "reasonstolive",
                 "rehabtherapy",
                 "sad",
                 "secondary_survivors",
                 "selfharm",
                 "siblingsupport",
                 "SMARTRecovery",
                 "socialanxiety",
                 "socialskills",
                 "StopSelfHarm",
                 "stopsmoking",
                 "SuicideWatch",
                 "uniqueminds",
                 "whatsbotheringyou",
                 "relationship_advice",
                 "howtonotgiveafuck",
                 "selfimprovement",
                 "personalfinance",
                 "confidence",
                 "love"
                 ]

data = {"subreddit": [], "content": []}

# Scraping hottest 200 posts for their title and the description
for i in subredditList:
    try:
        subreddit = reddit_read_only.subreddit(i)
        data["subreddit"].append(i)
        data["content"].append(subreddit.description)
        for post in subreddit.hot(limit=200):
            data["subreddit"].append(i)
            string = ""
            string += post.title
            string += " "
            string += post.selftext
            data["content"].append(string)
    except Exception:
        continue

# Converting it into a dataframe and saving to file
data = pd.DataFrame(data)
data.to_csv("data.csv")

# Reading file
data = pd.read_csv("data.csv")
le = preprocessing.LabelEncoder()

# Using label encoder for classification and removing stopwords and using stemmer
y_index = le.fit_transform(data.subreddit.values)
data['class'] = y_index
stemmer = PorterStemmer()
words = stopwords.words('english')
data['cleaned'] = data['content'].apply(lambda x: " ".join(
    [stemmer.stem(i) for i in re.sub("[^a-zA-Z]", " ", x).split() if i.lower() not in words and len(i) > 1]).lower())

# creating test and train dataset
x_train, x_test, y_train, y_test = train_test_split(data['cleaned'], data['class'], test_size=0.2)

# creating pipeline object to pipe TfidVectorizer result to select top 6000 features and using RCF
pipeline = Pipeline([('vect', TfidfVectorizer(ngram_range=(1, 2), stop_words="english", sublinear_tf=True)),
                     ('chi', SelectKBest(chi2, k=6000)),
                     ('clf', RandomForestClassifier(n_estimators=200, min_samples_split=20))])

# fitting test and train set and getting score of test set
model = pipeline.fit(x_train, y_train)
print("accuracy score: " + str(model.score(x_test, y_test)))

# fitting data with complete training data
model = pipeline.fit(data['cleaned'], data['class'])

# dumping model
from joblib import dump

dump(model, 'redditModel.joblib')

# making prediction based on the user's input
test = """
I broke off with my 5years long term relationship ex half a year ago as he was being mentally abusive to me, example if something doesn't goes his way he would shout at me although in his context he was just raising his voice. The relationship became very toxic and I decided to walk away from it as it was either break up or marriage if I were to continue staying on in this relationship. 

4months after the breakup he had zero contact with me leaving me very much heartbroken as I was still very much in love with him. I tried moving on by going on dates and meeting friends constantly keeping my mind occupied but My parents being kpo and not yet over the fact we broken up as i never told them the bad things he does to me and thinking we are making a mistake decided to invite him over for a family gathering one day. Surprised that he agreed to come it was the first time in 2 months since we last spoken, we had a serious conversation and found out that both of us was not over but both was too scar to take a step forward to do anything.... emotions came pouring out both of us feeling vulnerable and alcohol to blame we spent the night together. 

But after that night I realised we can never be together as old scars will never heal and it would be a mistake a try again. A month later my friend decided to tell me about his feeling for me. He had been there for me since the last year of my relationship with my ex and In a way he was a shoulder I could cry on and always lent his ear for me to rant since whenever I want to complain about things to my ex or was upset he would tell me Im overreacting and just playing the pity card ????, still being with my ex I didn't think anything much of his action aside from him being a good friend. After thinking about it for awhile I decided to give it a go and just try out this relationship.. 

Overtime I grew to really loving him and loving the way he treats me although he doesn't have much compared to my ex but atleast he treats me well... when everything seems to be going well I was hit with a news that u was pregnant. And it was definitely my ex as I have yet to done it with him. I told him about it and instead of being mad or asking me to abort it he hugged me and told me everything is going to be alright and if I ever decide to keep it he will raise it with me like his own If I want to abort he will be with me throughout the whole process.  

He asked whether do I want to tell my ex about it and what do I want to do.. part of me find that it's only right to tell him about this child yet a part of me thinks I shouldn't as I have no intention of raising it with him if I were to ever keep it, after all a child wouldn't just magically solve our problems. 

At the end of the day I decided to keep it with my boyfriend and to never let the child know about his real father but my plans got ruin when my ex found out about the kid and insist on marrying me and raising the child together saying it's his birth rights to do so as he is the biological father  and that things might change now he will change for me and the baby .. he wants to file a lawsuit if we don't agree and will do whatever legal means to gain rights over me and the child.  

Am I doing the right thing for my child? Should the child really grow up with the biological parents ? Should I give him another chance for the sake of the child..... even if it means gambling my happiness away.  Saying my boyfriend can never give the child everything that he can as he is a lawyer with a high paying income and my boyfriend just earning an average income. The child will be better off being able to choose to go to any school he wants and go and expensive vacations having all the toys he desire in the world and not being held back just because mummy and daddy can't afford... with all this drama going on I'm slowly sinking into depression and thinking maybe I should just get rid of the problem... if the child is gone..all this problems will be gone.... 

What should I do? I really hate myself for thinking like that.....
"""
y_predprob = model.predict_proba([test])
pred_To_Class_map = pd.DataFrame(y_predprob, columns=le.classes_)
pred_To_Class_map = pred_To_Class_map.transpose()
pred_To_Class_map = pred_To_Class_map.rename(columns={0: "prob"})
pred_To_Class_map = pred_To_Class_map.sort_values(by="prob", ascending=False)
recommendationList = [i for i in pred_To_Class_map.head(5).transpose()]

# collecting data of useful reddits
reddit_read_only = praw.Reddit(client_id="xnsIxViGuyQrTKs7VnnYCQ",  # your client id
                               client_secret="LuRt_v69n1GxxADiaLssD25fWOkuyw",  # your client secret
                               user_agent="safespace",
                               redirect_uri="http://localhost:8080")

urlAndInfoDict = {}

for i in recommendationList:
    subreddit = reddit_read_only.subreddit(i)
    urlAndInfoDict[subreddit.title] = [subreddit.public_description, subreddit.url, subreddit.subscribers]


