import { Injectable } from '@angular/core';
import { Router, CanActivate, CanActivateChild, CanDeactivate, CanLoad, Route, UrlSegment, ActivatedRouteSnapshot, RouterStateSnapshot, UrlTree } from '@angular/router';
import { Observable, of } from 'rxjs';
import { map, catchError } from 'rxjs/operators';
import { HttpClient } from '@angular/common/http';

import { AuthService } from '@services/auth.service';

@Injectable({
    providedIn: 'root'
})
export class AuthGuard implements CanActivate, CanActivateChild, CanDeactivate<unknown>, CanLoad {

    constructor (
        private authService: AuthService, 
        private router: Router,
        private httpClient: HttpClient,
    ) {}

    canActivate(
        route: ActivatedRouteSnapshot,
        state: RouterStateSnapshot): Observable<boolean | UrlTree> | Promise<boolean | UrlTree> | boolean | UrlTree {
        
        return this.checkLogin(state.url)
    }
    canActivateChild(
        childRoute: ActivatedRouteSnapshot,
        state: RouterStateSnapshot): Observable<boolean | UrlTree> | Promise<boolean | UrlTree> | boolean | UrlTree {
        return true;
    }
    canDeactivate(
        component: unknown,
        currentRoute: ActivatedRouteSnapshot,
        currentState: RouterStateSnapshot,
        nextState?: RouterStateSnapshot): Observable<boolean | UrlTree> | Promise<boolean | UrlTree> | boolean | UrlTree {
        return true;
    }
    canLoad(
        route: Route,
        segments: UrlSegment[]): Observable<boolean | UrlTree> | Promise<boolean | UrlTree> | boolean | UrlTree {
        const url = `/${route.path}`

        return this.checkLogin(url)
    }

    private checkLogin (url: string): Observable<boolean> {
        if (this.authService.isLoggedIn === undefined)
        {
            return this.httpClient.get<any>("/api/auth", {}).pipe(map((resp :any) => {
                this.authService.isLoggedIn = true

                if (url == '/login' || url == '/register') {
                    this.router.navigate(['/'])
                    return false
                }
                return true
            }),
            catchError((err: any) => {
                this.authService.isLoggedIn = false

                if (url == "/register") {
                    this.router.navigate([url])
                } else if (url != "/login") {
                    this.router.navigate(['/login'])            
                }
                return of (false)
            })
            )
        }

        if (url == '/login' || url == '/register') {
            if (this.authService.isLoggedIn) {
                this.router.navigate(['/'])
                return of(false)
            }
        }

        if (this.authService.isLoggedIn) {
            return of(true)
        }
        this.router.navigate(['/login'])
        return of(false)
    }
}
