import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;
import 'dart:convert';

class RegisterPage extends StatefulWidget {
  const RegisterPage({super.key});

  @override
  _RegisterPageState createState() => _RegisterPageState();
}

class _RegisterPageState extends State<RegisterPage> {
  final _formKey = GlobalKey<FormState>();
  String _username = '';
  String _email = '';
  String _password = '';

  Future<void> registerUser() async {
    final url = 'http://127.0.0.1:8080/api/register';

    final response = await http.post(
      Uri.parse(url),
      headers: {'Content-Type': 'application/json'},
      body: json.encode({
        'username': _username,
        'email': _email,
        'password': _password,
      }),
    );

    if (response.statusCode == 200 || response.statusCode == 201) {
      ScaffoldMessenger.of(context).showSnackBar(
        const SnackBar(content: Text('✅ Registration successful')),
      );
      Navigator.pushNamed(context, '/login');
    } else {
      ScaffoldMessenger.of(context).showSnackBar(
        SnackBar(
          content: Text('❌ Failed to register. (${response.statusCode})'),
        ),
      );
    }
  }

  //   @override
  //   Widget build(BuildContext context) {
  //     return Scaffold(
  //       appBar: AppBar(title: const Text('Register')),
  //       body: Padding(
  //         padding: const EdgeInsets.all(16.0),
  //         child: Form(
  //           key: _formKey,
  //           child: Column(
  //             children: [
  //               TextFormField(
  //                 decoration: const InputDecoration(labelText: 'Username'),
  //                 onChanged: (value) => setState(() => _username = value),
  //                 validator:
  //                     (value) =>
  //                         value!.isEmpty ? 'Please enter your username' : null,
  //               ),
  //               TextFormField(
  //                 decoration: const InputDecoration(labelText: 'Email'),
  //                 onChanged: (value) => setState(() => _email = value),
  //                 validator: (value) {
  //                   if (value!.isEmpty ||
  //                       !RegExp(r'\S+@\S+\.\S+').hasMatch(value)) {
  //                     return 'Please enter a valid email';
  //                   }
  //                   return null;
  //                 },
  //               ),
  //               TextFormField(
  //                 decoration: const InputDecoration(labelText: 'Password'),
  //                 obscureText: true,
  //                 onChanged: (value) => setState(() => _password = value),
  //                 validator:
  //                     (value) =>
  //                         value!.isEmpty ? 'Please enter a password' : null,
  //               ),
  //               const SizedBox(height: 20),
  //               ElevatedButton(
  //                 onPressed: () {
  //                   if (_formKey.currentState!.validate()) {
  //                     registerUser();
  //                   }
  //                 },
  //                 child: const Text('Register'),
  //               ),
  //             ],
  //           ),
  //         ),
  //       ),
  //     );
  //   }
  // }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Colors.grey[100],
      appBar: AppBar(
        title: const Text('Register'),
        backgroundColor: Colors.deepPurple,
        foregroundColor: Colors.white,
      ),
      body: Center(
        child: SingleChildScrollView(
          padding: const EdgeInsets.all(24.0),
          child: Column(
            children: [
              const SizedBox(height: 60),
              const Icon(
                Icons.person_add_alt_1,
                size: 80,
                color: Colors.deepPurple,
              ),
              const SizedBox(height: 20),
              const Text(
                "Create Your Account",
                style: TextStyle(
                  fontSize: 24,
                  fontWeight: FontWeight.bold,
                  color: Colors.deepPurple,
                ),
              ),
              const SizedBox(height: 40),
              Form(
                key: _formKey,
                child: Column(
                  children: [
                    TextFormField(
                      decoration: const InputDecoration(
                        labelText: 'Username',
                        border: OutlineInputBorder(),
                        prefixIcon: Icon(Icons.person),
                      ),
                      onChanged: (value) => setState(() => _username = value),
                      validator:
                          (value) =>
                              value!.isEmpty
                                  ? 'Please enter your username'
                                  : null,
                    ),
                    const SizedBox(height: 20),
                    TextFormField(
                      decoration: const InputDecoration(
                        labelText: 'Email',
                        border: OutlineInputBorder(),
                        prefixIcon: Icon(Icons.email),
                      ),
                      onChanged: (value) => setState(() => _email = value),
                      validator: (value) {
                        if (value!.isEmpty ||
                            !RegExp(r'\S+@\S+\.\S+').hasMatch(value)) {
                          return 'Please enter a valid email';
                        }
                        return null;
                      },
                    ),
                    const SizedBox(height: 20),
                    TextFormField(
                      decoration: const InputDecoration(
                        labelText: 'Password',
                        border: OutlineInputBorder(),
                        prefixIcon: Icon(Icons.lock),
                      ),
                      obscureText: true,
                      onChanged: (value) => setState(() => _password = value),
                      validator:
                          (value) =>
                              value!.isEmpty ? 'Please enter a password' : null,
                    ),
                    const SizedBox(height: 30),
                    SizedBox(
                      width: double.infinity,
                      child: ElevatedButton(
                        style: ElevatedButton.styleFrom(
                          backgroundColor: Colors.deepPurple,
                          foregroundColor: Colors.white,
                          minimumSize: const Size(200, 50),
                          shape: RoundedRectangleBorder(
                            borderRadius: BorderRadius.circular(12),
                          ),
                          elevation: 5,
                          shadowColor: Colors.black54,
                        ),
                        onPressed: () {
                          if (_formKey.currentState!.validate()) {
                            registerUser();
                          }
                        },
                        child: const Text(
                          'Register',
                          style: TextStyle(
                            fontSize: 18,
                            fontWeight: FontWeight.bold,
                            letterSpacing: 1.0,
                          ),
                        ),
                      ),
                    ),
                  ],
                ),
              ),
            ],
          ),
        ),
      ),
    );
  }
}
