{{define "content"}}
<form action="/users" method="post">
<center>
  <a href="https://www.amazon.pl/">
          <img alt="Qries" src="https://www.synointcdn.com/wp-content/uploads/2019/04/Amazon-Logo-PNG.png"
          width=160" height="70">
      </a>

      </center>

 <h2>Create New Password</h2><br>
<ul>
<li>usernames must contain only letters or numbers</li>
<li>usernames must be longer than 4 characters but shorter than 51</li>
<li>passwords must contain a uppercase letter, lowercase letter, number, and special character</li>
<li>passwords must be greater than 11 characters but less than 60</li>
</ul><br>
<form action="/forgotpwemailver{{.AuthInfo}}" method="POST">
<label for="password">Password:</label><br>
<input type="password" id="password" name="password"><br>
<label for="confirmpassword">Confirm Password:</label><br>
<input type="password" id="confirmpassword" name="confirmpassword"><br>
<input type="submit" value="Submit">
</form>
<br><br>
