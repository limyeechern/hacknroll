from joblib import load
import pandas as pd
from sklearn import preprocessing
from flask import Flask, request, jsonify
from flask_cors import CORS, cross_origin
from config import ApplicationConfig

df_y = pd.read_csv('./labels_500.csv', index_col=0)
model = load('./redditModel500.joblib')
redditData = pd.read_csv("./redditData.csv")
redditDict = {}
for i in redditData.values:
    redditDict[i[1]] = {"public_description": i[2], "url": i[3],
                        "subscribers": i[4]}
le = preprocessing.LabelEncoder()
le.fit(df_y['subreddit'].unique())

app = Flask(__name__)
CORS(app)
cors = CORS(app, resources={r'/*': {'origins': '*'}},
            supports_credentials=True)


@app.route("/getpredictions", methods=["POST"])
# @cross_origin(supports_credentials=True, resources={r'/*': {'origins': '*'}})
def userPrediction():
    input = request.json["data"]
    y_predprob = model.predict_proba([input])
    pred_To_Class_map = pd.DataFrame(y_predprob, columns=le.classes_)
    pred_To_Class_map = pred_To_Class_map.transpose()
    pred_To_Class_map = pred_To_Class_map.rename(columns={0: "prob"})
    pred_To_Class_map = pred_To_Class_map.sort_values(by="prob", ascending=False)
    recommendationList = [i for i in pred_To_Class_map.head(5).transpose()]
    urlAndInfoDict = []
    for i in recommendationList:
        publicDescription = redditDict[i]["public_description"]
        if type(publicDescription) is float:
            publicDescription = ""
        print("type is " + str(type(publicDescription)))
        urlAndInfoDict.append({"public_description": publicDescription,
                                   "url": "https://www.reddit.com" + redditDict[i]["url"],
                                   "subscribers": redditDict[i]["subscribers"],
                                   "title": redditDict[i]["url"]})

    print(urlAndInfoDict)
    return jsonify(urlAndInfoDict)


if __name__ == '__main__':
    print("test")
    app.run(host='0.0.0.0', port=8081)
