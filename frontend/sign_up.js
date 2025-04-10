document.addEventListener('DOMContentLoaded', function() {
  const form = document.getElementById('sign-up-form');
  if (form) {
    form.addEventListener('submit', function(event) {
      event.preventDefault();
      const username = document.getElementById('username').value;
      const password = document.getElementById('password').value;
      const confirmPassword = document.getElementById('confirm-password').value;
      const email = document.getElementById('email').value;

      if(password !== confirmPassword) {
        console.error("Passwords don't match");
        return;
      } 

      console.log(username, password, confirmPassword, email);
      createUser(username,password,email);
    });
  } else {
    console.error("Form not found");
  }
});

async function createUser(username, password, email) {
  const data = {
    "username": username,
    "password": password,
    "email": email
  };
  
  try {
    const response = await fetch('http://localhost:8080/createUser', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(data),
    });

    if (!response.ok) {
      throw new Error('Network response was not ok');
    }

    const result = await response.json();
    console.log(result);
  } catch (error) {
    console.error('Problem with creation:', error);
  }
}
