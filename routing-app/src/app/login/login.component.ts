import { Component, OnInit } from '@angular/core';
import { FormGroup, FormControl, Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { MessageService } from '@services/message.service';
import { AuthService } from '@services/auth.service';

@Component({
	selector: 'app-login',
	templateUrl: './login.component.html',
	styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {

	user = new FormGroup({
		email: new FormControl('', Validators.required),
		password: new FormControl('', Validators.required)
	})

	rememberme = new FormControl('')

	constructor(
		private httpClient: HttpClient,
		private messageService: MessageService,
		private authService: AuthService,
	) {}

	ngOnInit(): void {
		this.messageService.add('Load login page')
	}

	login()
	{
		this.messageService.add("Trying to login with data " + JSON.stringify(this.user.value))
		this.authService.login(JSON.stringify(this.user.value))
		console.log(this.messageService.messages)
	}

}
