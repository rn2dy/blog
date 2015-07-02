Why I want to do this? Here is the story - My current company uses ADP to log empolyee working hours, unfortunatelly, ADP portal uses old version of Java runtime. At the time of writing the portal only support JRE 6u17.

But most new computers has either JDK 7 (OpenJDK for Ubuntu) or above. It won't work!

Here is what is showing on my Ubuntu box (if you haven't already read the manual page for [http://linux.die.net/man/8/update-alternatives](update-alternative))

<pre class="prettyprint lang-sh">
\> update-alternatives --config java
There is only one alternative in link group java (providing /usr/bin/java): /usr/lib/jvm/java-7-openjdk-amd64/jre/bin/java
</pre>

*Tip: more information can be returned from `update-alternatives --query java`*

It is clear that I only have OpenJDK 7 installed (not the Oralce's official JDK).

Let's install JRE 6_u17 (required by the ADP portal) manually, if you want to install other versions same process applies.

**Step I**

Download jre-6u17 from here [http://www.oracle.com/technetwork/java/javase/downloads/java-archive-downloads-javase6-419409.html#jre-6u17-oth-JPR](Java SE 6 Archive)

For me I need to download *jre-6u17-linux-x64.bin*. (This is a excutable format which you can extract the content by just run it)

I also want to move it to where my default JDK installation is. Here is the sequence of commands I'd like to run.

<pre class="prettyprint lang-sh">
/> cd ~/Download
/> chmod +x jre-6u17-linux0-x64.bin && ./jre-6u17-linux-x64.bin
/> sudo mv jre1.6.0_17 /usr/lib/jvm
</pre>

**Step II**

Now is the time to use the [http://linux.die.net/man/8/update-alternatives](update-alternatives) command to make a alternative runtime for java:

<pre class="prettyprint lang-sh">
\> sudo update-alternatives --install /usr/bin/java java /usr/lib/jvm/jre1.6.0_17/bin/java 1
</pre>

To check if the alternative is really installed:

<pre class="prettyprint lang-sh">
\> update-alternatives --query java
Name: java
Link: /usr/bin/java
Slaves:
 java.1.gz /usr/share/man/man1/java.1.gz
Status: auto
Best: /usr/lib/jvm/java-7-openjdk-amd64/jre/bin/java
Value: /usr/lib/jvm/java-7-openjdk-amd64/jre/bin/java

Alternative: /usr/lib/jvm/java-7-openjdk-amd64/jre/bin/java
Priority: 1071
Slaves:
 java.1.gz /usr/lib/jvm/java-7-openjdk-amd64/jre/man/man1/java.1.gz

Alternative: /usr/lib/jvm/jre1.6.0_17/bin/java
Priority: 1
Slaves:
</pre>

Now to actually config it use the alternative (You might not have to if you set the priority to be higher then 1071)

<pre class="prettyprint lang-sh">
\> sudo update-alternatives --config java
There are 2 choices for the alternative java (providing /usr/bin/java).

  Selection    Path                                            Priority   Status
------------------------------------------------------------
* 0            /usr/lib/jvm/java-7-openjdk-amd64/jre/bin/java   1071      auto mode
  1            /usr/lib/jvm/java-7-openjdk-amd64/jre/bin/java   1071      manual mode
  2            /usr/lib/jvm/jre1.6.0_17/bin/java                1         manual mode

Press enter to keep the current choice[*], or type selection number:
</pre>

Choose 2 and enter and finally check Java version:


<pre class="prettyprint lang-sh">
java -version
java version "1.6.0_17"
Java(TM) SE Runtime Environment (build 1.6.0_17-b04)
Java HotSpot(TM) 64-Bit Server VM (build 14.3-b01, mixed mode)
</pre>

**Lastly**

I need to install firefox Java plugin (optional unless you are also stuck with ADP)

<pre class="prettyprint lang-sh">
\> sudo update-alternatives --install /usr/lib/mozilla/plugins/libjavaplugin.so mozilla-javaplugin.so /usr/lib/jvm/jre1.6.0_17/lib/amd64/libnpjp2.so 1
</pre>

It is probably not necessary to install plugin using update-alternative, you can just copy/move the .so file to the plugin directory.

After all this, I can finally log my hours and get paid. Life is tough, isn't it!
