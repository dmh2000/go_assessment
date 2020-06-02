/^=== RUN.*$/ { print "<div class=run>"$0"</div>"}
/^--- FAIL:.*$/ {print "<div class=fail>"$0"</div>"}
/.+_test/ {print "<div class=item>"$0"</div>"}