# Branching

résultats de la commande `git diff greet master` :

```sh
git diff greet master 
diff --git a/Makefile b/Makefile
index 12d4ffb..407082d 100644
--- a/Makefile
+++ b/Makefile
@@ -1,4 +1,3 @@
-# Ensure it runs the updated lib/hello.sh file
 TARGET="lib/hello.sh"
 
 run:
diff --git a/lib/greeter.sh b/lib/greeter.sh
deleted file mode 100644
index 2d14ca5..0000000
--- a/lib/greeter.sh
+++ /dev/null
@@ -1,6 +0,0 @@
-#!/bin/bash
-
-Greeter() {
-    who="$1"
-    echo "Hello, $who"
-}
diff --git a/lib/hello.sh b/lib/hello.sh
index 38c758b..8dde319 100644
--- a/lib/hello.sh
+++ b/lib/hello.sh
@@ -1,10 +1,6 @@
 #!/bin/bash
 
-source lib/greeter.sh
-
-name="$1"
-if [ -z "$name" ]; then
-    name="World"
-fi
-
-Greeter "$name"
+# Default is "World"
+# Author: Fabien OLIVIER
+name=${1:-"World"}
+echo "Hello, $name"
```

résultats de la commande `git diff master greet` :
```sh
diff --git a/Makefile b/Makefile
index 407082d..12d4ffb 100644
--- a/Makefile
+++ b/Makefile
@@ -1,3 +1,4 @@
+# Ensure it runs the updated lib/hello.sh file
 TARGET="lib/hello.sh"
 
 run:
diff --git a/lib/greeter.sh b/lib/greeter.sh
new file mode 100644
index 0000000..2d14ca5
--- /dev/null
+++ b/lib/greeter.sh
@@ -0,0 +1,6 @@
+#!/bin/bash
+
+Greeter() {
+    who="$1"
+    echo "Hello, $who"
+}
diff --git a/lib/hello.sh b/lib/hello.sh
index 8dde319..38c758b 100644
--- a/lib/hello.sh
+++ b/lib/hello.sh
@@ -1,6 +1,10 @@
 #!/bin/bash
 
-# Default is "World"
-# Author: Fabien OLIVIER
-name=${1:-"World"}
-echo "Hello, $name"
+source lib/greeter.sh
+
+name="$1"
+if [ -z "$name" ]; then
+    name="World"
+fi
+
+Greeter "$name"
```

résultats de la commande `git log --oneline --decorate --graph master..greet` :
```sh
* 5ab41a4 (origin/greet, greet) Makefile
* 8ccf32e change hello.sh in lib/
* 984bfc3 add file greeter.sh
```

résultats de la commande `git log --oneline --decorate --graph greet..master` :
```sh
* 874442b (HEAD -> master, origin/master) branching
```
