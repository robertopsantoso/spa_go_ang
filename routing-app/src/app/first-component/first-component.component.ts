import { Component, OnInit } from '@angular/core';
import { Router, ActivatedRoute, ParamMap } from '@angular/router';
import { Observable } from 'rxjs';
import { switchMap } from 'rxjs/operators';

// Services
import { MessageService } from '@services/message.service';

@Component({
	selector: 'app-first-component',
	templateUrl: './first-component.component.html',
	styleUrls: ['./first-component.component.css']
})
export class FirstComponentComponent implements OnInit {

	name: string;
	//heroes$: Observable<T>;
	selectedId: number;

	constructor(
		private route: ActivatedRoute,
		private router: Router,
		public messageService: MessageService
	) {}

	ngOnInit(): void {
		this.route.queryParams.subscribe(params => {
			this.name = params['name'];
			console.log(this.name);
			setTimeout(() => this.messageService.add('Console log name'), 10);
			
		})
		//this.route.paramMap.pipe(
		//	switchMap(params => {
		//		console.log(params);
		//		this.selectedId = Number(params.get('id'));
		//		console.log(this.selectedId);
		//		return null;
		//	})
		//)
		this.messageService.add('Load first component');
	}

	goToChildA() {
	  this.router.navigate(['child-a'], { relativeTo: this.route });
	}

}
