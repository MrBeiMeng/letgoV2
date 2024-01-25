## API 文档



{"query":"\n    query questionStats($titleSlug: String!) {\n  question(titleSlug: $titleSlug) {\n    stats\n  }\n}\n    ","variables":{"titleSlug":"median-of-two-sorted-arrays"},"operationName":"questionStats"}



resp:

{
    "data": {
        "question": {
            "stats": "{\"totalAccepted\": \"1.1M\", \"totalSubmission\": \"2.5M\", \"totalAcceptedRaw\": 1059264, \"totalSubmissionRaw\": 2527818, \"acRate\": \"41.9%\"}"
        }
    }
}





{"query":"\n    query singleQuestionTopicTags($titleSlug: String!) {\n  question(titleSlug: $titleSlug) {\n    topicTags {\n      name\n      slug\n      translatedName\n    }\n  }\n}\n    ","variables":{"titleSlug":"median-of-two-sorted-arrays"},"operationName":"singleQuestionTopicTags"}

resp:

{
    "data": {
        "question": {
            "topicTags": [
                {
                    "name": "Array",
                    "slug": "array",
                    "translatedName": "\u6570\u7ec4"
                },
                {
                    "name": "Binary Search",
                    "slug": "binary-search",
                    "translatedName": "\u4e8c\u5206\u67e5\u627e"
                },
                {
                    "name": "Divide and Conquer",
                    "slug": "divide-and-conquer",
                    "translatedName": "\u5206\u6cbb"
                }
            ]
        }
    }
}



{"query":"\n    query todayRecord {\n  todayRecord {\n    userStatus\n    comboTimes\n  }\n}\n    ","variables":{},"operationName":"todayRecord"}

resp:

{"data":{"todayRecord":[{"userStatus":"NOT_START","comboTimes":0}]}}



{"query":"\n    query questionTranslations($titleSlug: String!) {\n  question(titleSlug: $titleSlug) {\n    translatedTitle\n    translatedContent\n  }\n}\n    ","variables":{"titleSlug":"median-of-two-sorted-arrays"},"operationName":"questionTranslations"}

resp:

{
    "data": {
        "question": {
            "translatedTitle": "\u5bfb\u627e\u4e24\u4e2a\u6b63\u5e8f\u6570\u7ec4\u7684\u4e2d\u4f4d\u6570",
            "translatedContent": "<p>\u7ed9\u5b9a\u4e24\u4e2a\u5927\u5c0f\u5206\u522b\u4e3a <code>m</code> \u548c <code>n</code> \u7684\u6b63\u5e8f\uff08\u4ece\u5c0f\u5230\u5927\uff09\u6570\u7ec4&nbsp;<code>nums1</code> \u548c&nbsp;<code>nums2</code>\u3002\u8bf7\u4f60\u627e\u51fa\u5e76\u8fd4\u56de\u8fd9\u4e24\u4e2a\u6b63\u5e8f\u6570\u7ec4\u7684 <strong>\u4e2d\u4f4d\u6570</strong> \u3002</p>\n\n<p>\u7b97\u6cd5\u7684\u65f6\u95f4\u590d\u6742\u5ea6\u5e94\u8be5\u4e3a <code>O(log (m+n))</code> \u3002</p>\n\n<p>&nbsp;</p>\n\n<p><strong>\u793a\u4f8b 1\uff1a</strong></p>\n\n<pre>\n<strong>\u8f93\u5165\uff1a</strong>nums1 = [1,3], nums2 = [2]\n<strong>\u8f93\u51fa\uff1a</strong>2.00000\n<strong>\u89e3\u91ca\uff1a</strong>\u5408\u5e76\u6570\u7ec4 = [1,2,3] \uff0c\u4e2d\u4f4d\u6570 2\n</pre>\n\n<p><strong>\u793a\u4f8b 2\uff1a</strong></p>\n\n<pre>\n<strong>\u8f93\u5165\uff1a</strong>nums1 = [1,2], nums2 = [3,4]\n<strong>\u8f93\u51fa\uff1a</strong>2.50000\n<strong>\u89e3\u91ca\uff1a</strong>\u5408\u5e76\u6570\u7ec4 = [1,2,3,4] \uff0c\u4e2d\u4f4d\u6570 (2 + 3) / 2 = 2.5\n</pre>\n\n<p>&nbsp;</p>\n\n<p>&nbsp;</p>\n\n<p><strong>\u63d0\u793a\uff1a</strong></p>\n\n<ul>\n\t<li><code>nums1.length == m</code></li>\n\t<li><code>nums2.length == n</code></li>\n\t<li><code>0 &lt;= m &lt;= 1000</code></li>\n\t<li><code>0 &lt;= n &lt;= 1000</code></li>\n\t<li><code>1 &lt;= m + n &lt;= 2000</code></li>\n\t<li><code>-10<sup>6</sup> &lt;= nums1[i], nums2[i] &lt;= 10<sup>6</sup></code></li>\n</ul>\n"
        }
    }
}





{"query":"\n    query questionContent($titleSlug: String!) {\n  question(titleSlug: $titleSlug) {\n    content\n    editorType\n    mysqlSchemas\n    dataSchemas\n  }\n}\n    ","variables":{"titleSlug":"median-of-two-sorted-arrays"},"operationName":"questionContent"}

resp:

{
    "data": {
        "question": {
            "content": "<p>Given two sorted arrays <code>nums1</code> and <code>nums2</code> of size <code>m</code> and <code>n</code> respectively, return <strong>the median</strong> of the two sorted arrays.</p>\n\n<p>The overall run time complexity should be <code>O(log (m+n))</code>.</p>\n\n<p>&nbsp;</p>\n<p><strong class=\"example\">Example 1:</strong></p>\n\n<pre>\n<strong>Input:</strong> nums1 = [1,3], nums2 = [2]\n<strong>Output:</strong> 2.00000\n<strong>Explanation:</strong> merged array = [1,2,3] and median is 2.\n</pre>\n\n<p><strong class=\"example\">Example 2:</strong></p>\n\n<pre>\n<strong>Input:</strong> nums1 = [1,2], nums2 = [3,4]\n<strong>Output:</strong> 2.50000\n<strong>Explanation:</strong> merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.\n</pre>\n\n<p>&nbsp;</p>\n<p><strong>Constraints:</strong></p>\n\n<ul>\n\t<li><code>nums1.length == m</code></li>\n\t<li><code>nums2.length == n</code></li>\n\t<li><code>0 &lt;= m &lt;= 1000</code></li>\n\t<li><code>0 &lt;= n &lt;= 1000</code></li>\n\t<li><code>1 &lt;= m + n &lt;= 2000</code></li>\n\t<li><code>-10<sup>6</sup> &lt;= nums1[i], nums2[i] &lt;= 10<sup>6</sup></code></li>\n</ul>\n",
            "editorType": "CKEDITOR",
            "mysqlSchemas": [],
            "dataSchemas": []
        }
    }
}





{"query":"\n    query questionTitle($titleSlug: String!) {\n  question(titleSlug: $titleSlug) {\n    questionId\n    questionFrontendId\n    title\n    titleSlug\n    isPaidOnly\n    difficulty\n    likes\n    dislikes\n    categoryTitle\n  }\n}\n    ","variables":{"titleSlug":"median-of-two-sorted-arrays"},"operationName":"questionTitle"}

resp:

{
    "data": {
        "question": {
            "questionId": "4",
            "questionFrontendId": "4",
            "title": "Median of Two Sorted Arrays",
            "titleSlug": "median-of-two-sorted-arrays",
            "isPaidOnly": false,
            "difficulty": "Hard",
            "likes": 6979,
            "dislikes": 0,
            "categoryTitle": "Algorithms"
        }
    }
}