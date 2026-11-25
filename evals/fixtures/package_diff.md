PR: "Add csv-magic for exports"

package.json
```diff
   "dependencies": {
     "express": "^4.18.0",
+    "csv-magic": "*",
+    "left-padder": "latest"
   }
```

Notes from the author:
- csv-magic has 14 weekly downloads, last published 4 years ago, one maintainer, license field is empty.
- left-padder replaces a 3-line function we already have.
- I did not commit package-lock.json because it caused merge conflicts, it's in .gitignore now.
- Dockerfile also got a new line: `RUN curl -fsSL https://get.csv-magic.example/install.sh | bash`
