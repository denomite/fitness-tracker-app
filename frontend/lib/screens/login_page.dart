import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;
import 'dart:convert';
import 'package:shared_preferences/shared_preferences.dart';

class LoginScreen extends StatefulWidget {
  const LoginScreen({super.key});

  @override
  _LoginScreenState createState() => _LoginScreenState();
}

class _LoginScreenState extends State<LoginScreen> {
  final _formKey = GlobalKey<FormState>();
  String _email = '';
  String _password = '';
  bool _isLoading = false;

  @override
  void initState() {
    super.initState();
    print('Login screen initialized'); // Check if screen loads
  }

  Future<void> loginUser() async {
    print('Login button pressed'); // Confirm button triggers
    if (!_formKey.currentState!.validate()) {
      print('Form validation failed');
      setState(() => _isLoading = false); // Reset loading if validation fails
      return;
    }
    print('Form validation passed: email=$_email, password=$_password');

    setState(() => _isLoading = true);
    final url = 'http://192.168.1.11:8080/api/login';
    print('Attempting login to: $url');

    try {
      final response = await http.post(
        Uri.parse(url),
        headers: {'Content-Type': 'application/json'},
        body: json.encode({'email': _email, 'password': _password}),
      );
      print('Response: ${response.statusCode} - ${response.body}');

      if (response.statusCode == 200) {
        final data = jsonDecode(response.body);
        final token = data['token'];

        final prefs = await SharedPreferences.getInstance();
        await prefs.setString('token', token);

        if (mounted) {
          ScaffoldMessenger.of(
            context,
          ).showSnackBar(const SnackBar(content: Text("Login successful")));
          Navigator.pushReplacementNamed(context, '/home');
        }
      } else {
        final error = jsonDecode(response.body)['error'] ?? 'Unknown error';
        if (mounted) {
          ScaffoldMessenger.of(
            context,
          ).showSnackBar(SnackBar(content: Text('Login failed: $error')));
        }
      }
    } catch (e) {
      print('Login error: $e');
      if (mounted) {
        ScaffoldMessenger.of(
          context,
        ).showSnackBar(SnackBar(content: Text('Error: $e')));
      }
    } finally {
      if (mounted) setState(() => _isLoading = false);
    }
  }

  @override
  Widget build(BuildContext context) {
    print('Building LoginScreen'); // Track rebuilds
    return Scaffold(
      appBar: AppBar(title: const Text('Login')),
      body: Padding(
        padding: const EdgeInsets.all(16.0),
        child: Form(
          key: _formKey,
          child: Column(
            children: [
              TextFormField(
                decoration: const InputDecoration(labelText: 'Email'),
                validator:
                    (value) => value!.isEmpty ? 'Email is required' : null,
                onChanged: (value) {
                  setState(() {
                    _email = value; // Update state on change
                  });
                },
                initialValue: _email, // Preserve value on rebuild
              ),
              TextFormField(
                decoration: const InputDecoration(labelText: 'Password'),
                obscureText: true,
                validator:
                    (value) => value!.isEmpty ? 'Password is required' : null,
                onChanged: (value) {
                  setState(() {
                    _password = value; // Update state on change
                  });
                },
                initialValue: _password, // Preserve value on rebuild
              ),
              const SizedBox(height: 20),
              ElevatedButton(
                onPressed:
                    _isLoading
                        ? null
                        : () {
                          print(
                            'Button pressed, calling loginUser',
                          ); // Confirm press
                          loginUser();
                        },
                child:
                    _isLoading
                        ? const CircularProgressIndicator()
                        : const Text('Login'),
              ),
            ],
          ),
        ),
      ),
    );
  }
}

  // @override
  // Widget build(BuildContext context) {
  //   return Scaffold(
  //     backgroundColor: Colors.grey[100],
  //     appBar: AppBar(
  //       title: const Text("Login"),
  //       backgroundColor: Colors.deepPurple,
  //       foregroundColor: Colors.white,
  //     ),
  //     body: SingleChildScrollView(
  //       padding: const EdgeInsets.all(24.0),
  //       child: Column(
  //         children: [
  //           const SizedBox(height: 60),
  //           const Icon(Icons.lock_open, size: 80, color: Colors.deepPurple),
  //           const SizedBox(height: 20),
  //           const Text(
  //             "Welcome Back!",
  //             style: TextStyle(
  //               fontSize: 24,
  //               fontWeight: FontWeight.bold,
  //               color: Colors.deepPurple,
  //             ),
  //           ),
  //           const SizedBox(height: 40),
  //           Form(
  //             key: _formKey,
  //             child: Column(
  //               children: [
  //                 TextFormField(
  //                   decoration: const InputDecoration(
  //                     labelText: 'Email',
  //                     border: OutlineInputBorder(),
  //                     prefixIcon: Icon(Icons.email),
  //                   ),
  //                   onChanged: (value) => setState(() => _email = value),
  //                   validator:
  //                       (value) =>
  //                           value!.isEmpty ? 'Please enter your email' : null,
  //                 ),
  //                 const SizedBox(height: 20),
  //                 TextFormField(
  //                   decoration: const InputDecoration(
  //                     labelText: 'Password',
  //                     border: OutlineInputBorder(),
  //                     prefixIcon: Icon(Icons.lock),
  //                   ),
  //                   obscureText: true,
  //                   onChanged: (value) => setState(() => _password = value),
  //                   validator:
  //                       (value) =>
  //                           value!.isEmpty
  //                               ? 'Please enter your password'
  //                               : null,
  //                 ),
  //                 const SizedBox(height: 30),
  //                 SizedBox(
  //                   width: double.infinity,
  //                   child: ElevatedButton(
  //                     style: ElevatedButton.styleFrom(
  //                       backgroundColor: Colors.deepPurple, // Vibrant color
  //                       foregroundColor: Colors.white, // Text color
  //                       minimumSize: const Size(200, 50), // Size of the button
  //                       shape: RoundedRectangleBorder(
  //                         borderRadius: BorderRadius.circular(12),
  //                       ),
  //                       elevation: 5,
  //                       shadowColor: Colors.black54,
  //                     ),
  //                     onPressed: () => Navigator.pushNamed(context, '/login'),
  //                     child: const Text(
  //                       "Login",
  //                       style: TextStyle(
  //                         fontSize: 18,
  //                         fontWeight: FontWeight.bold,
  //                         letterSpacing: 1.0,
  //                       ),
  //                     ),
  //                   ),
  //                 ),
  //               ],
  //             ),
  //           ),
  //         ],
  //       ),
  //     ),
  //   );
  // }

