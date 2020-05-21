# Hasagi project

## Thuật ngữ
- `Parent`: Phân quyền cha mẹ, người quản lí `child`
- `Child`: Phân quyền con, người chịu sự quản lí của cha mẹ
- `Alert location`: Địa điểm mà khi user đến đó thì sẽ thông báo cho người quản lí của user
- `Location`: Vị trí của child

## Chức năng
### Parent

- [x] Tạo alert location cho child
- [ ] Xóa alert location
- [ ] Sửa alert location
- [ ] Xem thông tin một (danh sách) alert location
<hr>

- [ ] Lấy vị trí hiện tại của child

### Child
- [x] Update location kèm thông báo đến `parent` nếu đó là alert location

## API
Root: `https://hasagi1998.herokuapp.com/api/v1/`

### Tạo alert cho child

Method: POST
1. Header

|key|Type|Sample value
|---|---|---
|Authorization|string| Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJhdWQiOiIyIiwianRpIjoiMjdhYjFlNTJlZTQxMDZkMDQxYzJlYzQwNGU2MWVkMjY1NTA1ZjA5NjA4MTI1MzNhN2U5YjY5MWEwNjdiZWFjMmZmMzAwYWZmYWM1MTdjYWUiLCJpYXQiOjE1OTAwMjgyNTUsIm5iZiI6MTU5MDAyODI1NSwiZXhwIjoxNjIxNTY0MjU1LCJzdWIiOiI0Iiwic2NvcGVzIjpbXX0.VCj1DkWDS6rlhkpDP8xhfGXFD12_xANCj9vg7hLcGlZ9pmQCRPXKtdZFzbxbEAYGo5TXOb7jj1eOeL-I0bnz2Ok1XNfnQiFA28fa4i4hcnie_zsF23GMjLZ82h0mhi22xspJQBQ1tDujiQcxm7pjA3QgkNecoF4D9nhiuJZvU9ma5IFIhvhkbNs9Z_U0FR--raMZLJzR82G4d2r1Rjq2SPckTVY6r3227Q_GEEdF9kNW1xhQ3M90Zgb1SWarA_MvE4mD4Vd5TAUxOzP7gREfTBb5g6gEtMHOZ4gK-uuMWRLUOFPx4HV__nuqoOfm4lkL_AaLgg6N4gH6dsHD5siS5GMAtw3reYfOz1SVta-Lmg5JRTNe0zQsN6lolXKfMCaV4w-R77wKvVnTxACWn9HQXQNZxIgd-TOJhKj2RgIMKxhoHluW_rVLGqCa_qauW2hRpP1W23SjRA4x-onL81GDW9FU19nuzJ89m1a_g-m0es08XeF3lxmG-ccnoWaTgEWht_TeA7D5aT_bZXbvSSv_qYfZ3vuQmtlghWaqXFnAMgrougPmXhUD3b8Bz33gix-5YW6s2NrpfM3hsyVSaC-tbMyDnhiR6fp5AVtaKXcYDKGD1QNXbaj_8myjnenueXL6vgVQjryQYXBmXCw6htf6vnX0D8vXlW8-m6cz17WUjd0
|Content-Type|string|application/json

2. Body

|key|Type|Sample value
|---|---|---
|longitude|float|156.234
|latitude|float|123.456
|name|string|Trường học của mỡ
|userid|uint|1 (userID of child)
