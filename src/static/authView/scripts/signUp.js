document.addEventListener('DOMContentLoaded', function() {
handleSignUpForm();  
});

function handleSignUpForm(){
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

      createUser(username,password,email);
    });
  } else {
    console.error("Form not found");
  }
}

async function createUser(username, password, email) {
  const data = {
    "username": username,
    "password": password,
    "email": email
  };
  
  try {
    const response = await fetch('/api/createUser', {
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
    alert('Account created.');
    window.location.href = result.redirectTo;

  } catch (error) {
    console.error('Problem with creation:', error);
  }
}
