import { Component, OnInit } from '@angular/core';
import { Router, ActivatedRoute, ParamMap } from '@angular/router';
import { Observable } from 'rxjs';

@Component({
	selector: 'app-child-a',
	templateUrl: './child-a.component.html',
	styleUrls: ['./child-a.component.css']
})
export class ChildAComponent implements OnInit {

	constructor(
		private route: ActivatedRoute,
		private router: Router
	) { }

	ngOnInit(): void {
		const heroId = this.route.snapshot.paramMap.get('id');
		console.log(heroId);
	}

}
