
<hr style="background:#ffd04c;margin: 0 200px;height:18px;border-radius:5px">


# 无重复字符的最长子串

[<span style="font-weight:bold;font-size:14px">中文 👈</span>](./README-zh.md) | [<span style="font-weight:bold;font-size:14px">ENGLISH</span>](./README-en.md)

<span style="font-weight:bold;font-size:14px">难度</span> <span style="background:#ffa400;border-radius:5px;padding:1px 5px;font-weight:bold;color:#ffffff">中等</span>   <span style="font-weight:bold;font-size:14px">地址</span> [3.无重复字符的最长子串](https://leetcode.cn/problems/longest-substring-without-repeating-characters)

<p>给定一个字符串 <code>s</code> ，请你找出其中不含有重复字符的&nbsp;<strong>最长子串&nbsp;</strong>的长度。</p>

<p>&nbsp;</p>

<p><strong>示例&nbsp;1:</strong></p>

<pre>
<strong>输入: </strong>s = "abcabcbb"
<strong>输出: </strong>3 
<strong>解释:</strong> 因为无重复字符的最长子串是 <code>"abc"</code>，所以其长度为 3。
</pre>

<p><strong>示例 2:</strong></p>

<pre>
<strong>输入: </strong>s = "bbbbb"
<strong>输出: </strong>1
<strong>解释: </strong>因为无重复字符的最长子串是 <code>"b"</code>，所以其长度为 1。
</pre>

<p><strong>示例 3:</strong></p>

<pre>
<strong>输入: </strong>s = "pwwkew"
<strong>输出: </strong>3
<strong>解释: </strong>因为无重复字符的最长子串是&nbsp;<code>"wke"</code>，所以其长度为 3。
&nbsp;    请注意，你的答案必须是 <strong>子串 </strong>的长度，<code>"pwke"</code>&nbsp;是一个<em>子序列，</em>不是子串。
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>0 &lt;= s.length &lt;= 5 * 10<sup>4</sup></code></li>
	<li><code>s</code>&nbsp;由英文字母、数字、符号和空格组成</li>
</ul>


<hr style="background:#ffd04c;margin: 0 60px">


<span style="font-weight:bold;font-size:14px">通过率：39.5%</span>  <span style="font-weight:bold;font-size:14px" alt="哈希表|字符串|滑动窗口">点击查看标签</span>

